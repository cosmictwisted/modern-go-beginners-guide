package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       string
}

func readCSV(fileName string) {
	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	for {
		line, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for item := range line {
			fmt.Printf("%s\n", line[item])
		}

	}
}

func readAllCSV(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(f)

	// Skip the first header line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	data, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return data, nil
}

func printFriends(data [][]string, err error) {
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range data {
		fmt.Printf("Name: %s %s, Email: %s, Age: %s\n", d[0], d[1], d[2], d[3])
	}
}

func parseAllCSV(data [][]string, err error) {
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range data {
		u := User{
			FirstName: d[0],
			LastName:  d[1],
			Email:     d[2],
			Age:       d[3],
		}
		fmt.Printf("%s %s is %s years old.\n", u.FirstName, u.LastName, u.Age)
	}
}

// Write to a csv file
func newFriends(friends [][]string) {
	f, err := os.Create("new_friends.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, friend := range friends {
		if err := w.Write(friend); err != nil {
			log.Fatalln("error writing to file", err)
		}
	}
}

func main() {
	readCSV("numbers.csv")
	printFriends(readAllCSV("friends.csv"))
	parseAllCSV(readAllCSV("friends.csv"))

	friends := [][]string{
		{"Rob", "Pike", "rob@example.com", "51"},
		{"Jim", "Carrey", "jim@example.com", "55"},
	}
	newFriends(friends)
}
