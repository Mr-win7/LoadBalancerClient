package main

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// Backend Contains ip behind dns with its status.
type Backend struct {
	IP      string
	isAlive bool
	mutex   sync.RWMutex
}

// GetIsAlive Check is backend alive.
func (backend *Backend) GetIsAlive() bool {
	backend.mutex.RLock()
	result := backend.isAlive
	backend.mutex.RUnlock()
	return result
}

// SetIsAlive Change the status of backend.
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

// IPListFetch Get all available ip from dns.
// Result is saved in global variable IpList.
func IPListFetch(domain string) {
	cmd := exec.Command("dig", "+noall", "+answer", domain)
	output, err := cmd.Output()
	if err != nil {
		return
	}
	ParseDigResponse(string(output))
}
