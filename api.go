package main

import (
	"encoding/json"
	"net/http"
)

// CreateTask
func CreateTask(w http.ResponseWriter, r *http.Request) {

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully!"})
}

// get task through description
func GetTask(w http.ResponseWriter, r *http.Request) {
	description := r.URL.Query().Get("description")
	if description == "" {
		http.Error(w, "Please provide a task description", http.StatusBadRequest)
		return
	}

	for _, t := range tasks {
		if t.Description == description {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

//update

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	for i := range tasks {
		if tasks[i].Description == updatedTask.Description {

			if tasks[i].OwnerId != updatedTask.OwnerId {
				http.Error(w, "Unauthorized: only the owner can update this task", http.StatusForbidden)
				return
			}

			tasks[i].IsCompleted = updatedTask.IsCompleted

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully!"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

//delete

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	description := r.URL.Query().Get("description")
	if description == "" {
		http.Error(w, "Please provide a task description", http.StatusBadRequest)
		return
	}

	for i, t := range tasks {
		if t.Description == description {

			tasks = append(tasks[:i], tasks[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully!"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
