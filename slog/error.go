package slog

import (
	"log"
	"os"
)

var errLog *log.Logger

func init() {
	errLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func ErrorPrint(args ...interface{}) {
	errLog.Println(args...)
}
