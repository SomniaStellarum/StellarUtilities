package slog

import (
	"io/ioutil"
	"log"
	"os"
)

var debug *log.Logger

func init() {
	debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime)
}

func SetDebug() {
	debug.SetOutput(os.Stdout)
}

func DebugPrint(args ...interface{}) {
	debug.Println(args...)
}
