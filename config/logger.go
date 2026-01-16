package config

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("blog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("failed to open log file: %v", err)
	}

	Logger = log.New(file, "BLOG", log.Ldate|log.Ltime|log.Lshortfile)
}
