package task

import "sync"

type Task struct {
	Description string `json:"description"`
	OwnerId     int    `json:"owner_id"`
	IsCompleted bool   `json:"is_completed"`
}

var (
	Tasks []Task
	Mutex sync.RWMutex
)
