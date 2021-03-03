package main

import (
	"os/exec"
)

// GenerateCommand creates a command object from the given command and arguments
func GenerateCommand(cmd string, args ...string) (command *exec.Cmd) {
	return exec.Command(cmd, args...)
}

// Execute runs a command object and returns the command output
func Execute(cmd *exec.Cmd) (output []byte, err error) {
	output, err = cmd.CombinedOutput()
	return
}