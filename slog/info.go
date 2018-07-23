package slog

import (
	"log"
	"os"
)

var info *log.Logger

func init() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoPrint(args... interface{}) {
	info.Println(args...)
}