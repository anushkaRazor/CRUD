package api

import (
	"encoding/json"
	"net/http"

	"github.com/anushkaRazor/CRUD/internal/logger"
	"github.com/anushkaRazor/CRUD/internal/task"
)

func respJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}


func fetchTask(w http.ResponseWriter, id, user string, checkOwner bool) (task.Task, bool) {
	task.Mutex.RLock()
	t, ok := task.Tasks[id]
	task.Mutex.RUnlock()

	if !ok {
		logger.L.Printf("Task not found: ID=%s", id)
		http.Error(w, "Task not found", http.StatusNotFound)
		return task.Task{}, false
	}

	if checkOwner && t.OwnerId != user {
		logger.L.Printf("Unauthorized: task=%+v user=%s", t, user)
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return task.Task{}, false
	}

	logger.L.Printf("Task fetched: %+v", t)
	return t, true
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		logger.L.Printf("Create failed: invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task.Mutex.Lock()
	task.Tasks[t.ID] = t
	task.Mutex.Unlock()

	logger.L.Printf("Task created: %+v", t)
	respJSON(w, map[string]string{"msg": "Created"}, http.StatusOK)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		logger.L.Println("Get failed: missing ID")
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	t, ok := fetchTask(w, id, "", false)
	if !ok {
		return
	}
	respJSON(w, t, http.StatusOK)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-ID")
	if user == "" {
		logger.L.Println("Update failed: missing user ID")
		http.Error(w, "Missing user ID", http.StatusUnauthorized)
		return
	}

	var upd task.Task
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		logger.L.Printf("Update failed: invalid JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	t, ok := fetchTask(w, upd.ID, user, true)
	if !ok {
		return
	}

	task.Mutex.Lock()
	t.Description = upd.Description
	t.IsCompleted = upd.IsCompleted
	task.Tasks[upd.ID] = t
	task.Mutex.Unlock()

	logger.L.Printf("Task updated: %+v by user=%s", t, user)
	respJSON(w, map[string]string{"msg": "Updated"}, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := r.Header.Get("User-ID")
	if id == "" || user == "" {
		logger.L.Printf("Delete failed: missing ID or user (ID=%s, User=%s)", id, user)
		http.Error(w, "ID and user required", http.StatusBadRequest)
		return
	}

	t, ok := fetchTask(w, id, user, true)
	if !ok {
		return
	}

	task.Mutex.Lock()
	delete(task.Tasks, id)
	task.Mutex.Unlock()

	logger.L.Printf("Task deleted: %+v by user=%s", t, user)
	respJSON(w, map[string]string{"msg": "Deleted"}, http.StatusOK)
}

func Health(w http.ResponseWriter, r *http.Request) {
	logger.L.Println("Status healthy")
	respJSON(w, map[string]string{"status": "ok"}, http.StatusOK)
}
