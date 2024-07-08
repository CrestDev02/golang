package logger

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "myapp: ", log.LstdFlags|log.Lshortfile)
}
