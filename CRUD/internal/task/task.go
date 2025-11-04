package task

import "sync"

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	OwnerId     string `json:"ownerId"`
}

var (
	Tasks = make(map[string]Task) 
	Mutex sync.RWMutex
)
