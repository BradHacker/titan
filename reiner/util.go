package main

import "os"

// GetProcessID returns the current PID of the Beacon
func GetProcessID() (pid int) {
	pid = os.Getpid()
	return
}