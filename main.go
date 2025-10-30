package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	Log()
	http.HandleFunc("/create", CreateTask)
	http.HandleFunc("/read", GetTask)
	http.HandleFunc("/update", UpdateTask)
	http.HandleFunc("/delete", DeleteTask)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
