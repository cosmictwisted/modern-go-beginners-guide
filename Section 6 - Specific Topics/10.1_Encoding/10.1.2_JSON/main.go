package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type Data map[string]interface{}

func readJSON(fileName string) ([]User, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var friends []User

	err = json.Unmarshal(file, &friends)
	if err != nil {
		log.Fatalln(err)
	}

	return friends, nil
}

func readUnstructuredJSON(fileName string) ([]Data, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var d []Data

	err = json.Unmarshal(file, &d)
	if err != nil {
		log.Fatalln(err)
	}

	return d, nil
}

func encodeJSON(friends []User) {
	data, err := json.Marshal(friends)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func encodePrettyJSON(friends []User) {
	data, err := json.MarshalIndent(friends, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func main() {
	friends, err := readJSON("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	for _, friend := range friends {
		fmt.Println(friend)
	}

	encodeJSON(friends)
	encodePrettyJSON(friends)

	data, err := readUnstructuredJSON("unstructured.json")
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range data {
		for k, v := range d {
			fmt.Printf("Key: %s\tValue: %v\n", k, v)
		}
		fmt.Println()
	}
}
