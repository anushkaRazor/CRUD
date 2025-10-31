package main

import (
	"fmt"
	"net/http"
)

func main() {

	Log()
	logger.Println("Logger initialized")

	http.HandleFunc("/create", CreateTask)
	http.HandleFunc("/read", GetTask)
	http.HandleFunc("/update", UpdateTask)
	http.HandleFunc("/delete", DeleteTask)
	http.HandleFunc("/ping", HealthCheck)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}