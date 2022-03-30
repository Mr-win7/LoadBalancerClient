package main

import (
	"log"
	"os"
)

var (
  // INFO logger for info level
	INFO    *log.Logger
  // WARNING logger for warning level
	WARNING *log.Logger
  // ERROR logger for error level
	ERROR   *log.Logger
)

// InitLog Init logger.
func InitLog() {
	INFO = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	WARNING = log.New(os.Stdout, "WARNING: ", log.LstdFlags)
	ERROR = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
}
