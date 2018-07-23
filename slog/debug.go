package slog

import (
	"log"
	"io/ioutil"
	"os"
)

var debug *log.Logger

func init() {
	debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func SetDebug() {
	debug.SetOutput(os.Stdout)
}

func DebugPrint(args... interface{}) {
	debug.Println(args...)
}