package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 2500

func main() {
	rootCtx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(rootCtx, time.Duration(timeout)*time.Millisecond)
	defer cancel()

	user, err := getUser(ctxWithTimeout, 3)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Result: ", user)
}

func getUser(ctx context.Context, userid int) (string, error) {
	result := make(chan string)
	go func() {
		result <- slowDBQuery(userid)
		close(result)
	}()

	for {
		select {
		case user := <-result:
			return user, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func slowDBQuery(userid int) string {
	rand.Seed(time.Now().Unix())
	delay := rand.Intn(5)
	fmt.Println("Delay:", delay)
	fmt.Printf("Searching for userid: %d\n", userid)
	time.Sleep(time.Duration(delay) * time.Second)

	return fmt.Sprintf("Found user with userid: %d", userid)
}
