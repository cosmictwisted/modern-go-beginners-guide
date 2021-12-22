package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Post struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	now := time.Now()
	fmt.Println("Datetime:", now)

	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	after := time.Since(now)
	fmt.Println("Elapsed:", after)

	p := &Post{
		Body:      "Humpty Dumpty sat on a wall",
		CreatedAt: time.Now(),
	}

	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
