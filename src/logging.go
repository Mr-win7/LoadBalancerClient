package main

import (
  "log"
  "os"
)

var (
  INFO *log.Logger
  WARNING *log.Logger
  ERROR *log.Logger
)

func InitLog() {
  INFO = log.New(os.Stdout, "INFO: ", log.LstdFlags)
  WARNING = log.New(os.Stdout, "WARNING: ", log.LstdFlags)
  ERROR = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
}
