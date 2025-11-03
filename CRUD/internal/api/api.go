package api

import (
	"encoding/json"
	"net/http"
	"github.com/anushkaRazor/CRUD/internal/logger"
	"github.com/anushkaRazor/CRUD/internal/task"
)

// CreateTask
func CreateTask(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Println("Create Task endpoint hit")

	var newTask task.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	task.Mutex.Lock()
	task.Tasks = append(task.Tasks, newTask)
	task.Mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully!"})
}

// get task
func GetTask(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Println("Read task endpoint hit")
	description := r.URL.Query().Get("description")
	if description == "" {
		http.Error(w, "Please provide a task description", http.StatusBadRequest)
		return
	}

	task.Mutex.RLock()
	defer task.Mutex.RUnlock()
	for _, t := range task.Tasks {
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

	logger.Logger.Println("Update task endpoint hit")
	var updatedTask task.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	task.Mutex.Lock()
	defer task.Mutex.Unlock()

	for i := range task.Tasks {
		if task.Tasks[i].Description == updatedTask.Description {

			if task.Tasks[i].OwnerId != updatedTask.OwnerId {
				http.Error(w, "Unauthorized: only the owner can update this task", http.StatusForbidden)
				return
			}

			task.Tasks[i].IsCompleted = updatedTask.IsCompleted

			w.Header().Set("Content-Type", "application/json")git
			json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully!"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

//delete

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Println("Delete task endpoint hit")
	var deletedTask task.Task

	description := r.URL.Query().Get("description")
	if description == "" {
		http.Error(w, "Please provide a task description", http.StatusBadRequest)
		return
	}

	task.Mutex.Lock()
	defer task.Mutex.Unlock()

	for i := range task.Tasks {
		if task.Tasks[i].Description == description {

			if task.Tasks[i].OwnerId != deletedTask.OwnerId {
				http.Error(w, "Unauthorized: only the owner can delete this task", http.StatusForbidden)
				return
			}

			task.Tasks = append(task.Tasks[:i], task.Tasks[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully!"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// healthCheck

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	logger.Logger.Println("HealthCheck task endpoint hit")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Serivce is healthy\n"))

}