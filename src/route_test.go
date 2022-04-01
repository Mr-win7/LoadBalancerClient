package main

import (
  "testing"
  "sync"
)

func TestGetNextIndex(t *testing.T) {
  backendList = []Backend{{"104.71.49.125", true, sync.RWMutex{}}}
  currentIndex = 0
  t.Logf("len(backendList): %d", len(backendList))
  expected := uint64(0)
  actual, err := getNextIndex()
  if expected != actual || expected != currentIndex || err != nil {
    t.Errorf("expected: %d, actual: %d", expected, actual)
  }

}

func TestGetRouteRecord(t *testing.T) {
  backendList = []Backend{}
  currentIndex = 0
  t.Logf("len(backendList): %d", len(backendList))
  backend, err := getRouteRecord()
  if err == nil || backend != nil {
    t.Errorf("get invalid backend.")
  }
}
