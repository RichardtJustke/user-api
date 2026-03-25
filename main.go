package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// rota teste
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	fmt.Println("Servidor rodando em http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
