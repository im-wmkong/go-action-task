package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully opened users.json")
	users := new(Users)
	err = json.Unmarshal(data, users)
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users.Users {
		fmt.Println("User Type: ", user.Type)
		fmt.Println("User Age: ", user.Age)
		fmt.Println("User Name: ", user.Name)
		fmt.Println("Facebook Url: ", user.Social.Facebook)
	}
	fmt.Println(users)
}
