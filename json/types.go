package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID      string  `json:"id"`
	Profile Profile `json:"profile,omitempty"`
	Karma   int     `json:"karma,omitempty"`
}

type Profile struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// test json omitempty
func main() {
	user := User{
		ID: "test",
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(bytes)
}
