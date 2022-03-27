package main

import (
  "bufio"
  "strings"
  "os/exec"
  "regexp"
  "sync"
)

type Backend struct {
  Ip string
  isAlive bool
  mutex sync.RWMutex
}

func (backend *Backend) GetIsAlive() bool {
  backend.mutex.RLock()
  result := backend.isAlive
  backend.mutex.RUnlock()
  return result
}

func (backend *Backend) SetIsAlive(isAlive bool) {
  backend.mutex.Lock()
  backend.isAlive = isAlive
  backend.mutex.Unlock()
}

var backendList []Backend

// ParseDigResponse Get backend list from dig output.
// Only preserve A records.
func ParseDigResponse(response string) {
  scanner := bufio.NewScanner(strings.NewReader(response))
  pattern, _ := regexp.Compile(`\s+`)
  for scanner.Scan() {
    line := scanner.Text()
    recordSlice := pattern.Split(line, -1) // all sequence
    recordType := recordSlice[3]

    if recordType != "A" {
      continue
    }

    ipAddress := recordSlice[4]
    backendList = append(backendList, Backend{ipAddress, true, sync.RWMutex{}})
  }
}

// IpListFetch Get all available ip from dns.
// Result is saved in global variable IpList.
func IpListFetch(domain string) {
  cmd := exec.Command("dig", "+noall", "+answer", domain)
  output, err := cmd.Output()
  if err != nil {
    return;
  }
  ParseDigResponse(string(output))
}

