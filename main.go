package main

import (
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("Hello world")
}
