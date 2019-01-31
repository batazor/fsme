package main

import (
	"fmt"
	"github.com/batazor/fsme/fsm"
)

func main() {
	// Create new FSM
	machine, err := fsm.New()

	// Add rule
	machine.AddStateTransitionRules("a", "b", "c")
	machine.AddStateTransitionRules("b", "d", "e")
	machine.AddStateTransitionRules("c", "k")
	machine.AddStateTransitionRules("d", "a")
	machine.AddStateTransitionRules("e", "k")
	machine.AddStateTransitionRules("k")

	// Init State
	err = machine.SetStateTransition("b")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Next step
	err = machine.SetStateTransition("e")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Next step
	err = machine.SetStateTransition("k")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("State: %s\n", machine.GetCurrentState())
}
