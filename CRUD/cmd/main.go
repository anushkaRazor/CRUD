package main

import (
	"fmt"
	"net/http"

	"github.com/anushkaRazor/CRUD/internal/api"
	"github.com/anushkaRazor/CRUD/internal/logger"
)

func main() {
	logger.Log()
	fmt.Println("Logger initialized")

	http.HandleFunc("/create", api.CreateTask)
	http.HandleFunc("/read", api.GetTask)
	http.HandleFunc("/update", api.UpdateTask)
	http.HandleFunc("/delete", api.DeleteTask)
	http.HandleFunc("/ping", api.HealthCheck)

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the CRUD App!"))

	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
