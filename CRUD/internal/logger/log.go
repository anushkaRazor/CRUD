package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var L *log.Logger

func Log() {

	logDir := "log"
	timestamp := time.Now().Format("20060102_150405")
	FileName := fmt.Sprintf("log_file_%s.txt", timestamp)
	filePath := filepath.Join(logDir, FileName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Sync()
	if err != nil {
		log.Fatalf("Error opening log file:%v", err)
	}

	L = log.New(file, "TASK-LOG : ", log.Ldate|log.Ltime|log.Lshortfile)
	L.Println("Log initialized successfully")
}