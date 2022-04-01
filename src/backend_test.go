package main

import (
  "testing"
  "sync"
  "reflect"
)

func TestGetIsAlive(t *testing.T) {
  backend := Backend{"1.1.1.1", false, sync.RWMutex{}}
  actual := backend.GetIsAlive()
  expected := false
  if expected != actual {
    t.Errorf("expected: %t, actual: %t", expected, actual)
  }

  backend.isAlive = true
  actual = backend.GetIsAlive()
  expected = true
  if expected != actual {
    t.Errorf("expected: %t, actual: %t", expected, actual)
  }
}

func TestSetIsAlive(t *testing.T) {
  backend := Backend{"1.1.1.1", false, sync.RWMutex{}}
  backend.SetIsAlive(true)
  actual := backend.GetIsAlive()
  expected := true
  if expected != actual {
    t.Errorf("expected: %t, actual: %t", expected, actual)
  }
  backend.SetIsAlive(false)
  expected = false
  actual = backend.GetIsAlive()
  if expected != actual {
    t.Errorf("expected: %t, actual: %t", expected, actual)
  }
}

func TestParseDigResponse(t *testing.T) {
  input := ""
  ParseDigResponse(input)
  expected := make([]Backend, 0)
  actual := backendList
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("expected: %#v, actual: %#v", expected, actual)
  }

  input = "e9428.a.akamaiedge.net.	30	IN	A	184.51.137.48"
  ParseDigResponse(input)
  expected = []Backend{{"184.51.137.48", true, sync.RWMutex{}}}
  actual = backendList
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("expected: %v, actual: %v", expected, actual)
  }
}
