// Ref - https://www.youtube.com/watch?v=JlmYLPxwVzQ
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type token struct {
	mu   sync.RWMutex
	data map[string]int
}

var tokens = &token{data: make(map[string]int)}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go save(fmt.Sprintf("token_id_%d", rand.Intn(100)), rand.Intn(100), wg)
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			tokens.mu.RLock()
			fmt.Println(tokens.data)
			tokens.mu.RUnlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(tokens.data)
}

func save(k string, v int, wg *sync.WaitGroup) {
	tokens.mu.Lock()
	tokens.data[k] = v
	tokens.mu.Unlock()
	wg.Done()
}

func saveSlow(k string, v int) {
	tokens.data[k] = v
}
