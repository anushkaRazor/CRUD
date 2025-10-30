package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var logger *log.Logger

func Log() {

	logDir := "log"
	timestamp := time.Now().Format("20060102_150405")
	FileName := fmt.Sprintf("log_file_%s.txt", timestamp)
	filePath := filepath.Join(logDir, FileName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file:%v", err)
	}

	logger = log.New(file, "TASK-LOG : ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Log initialized successfully")
}
