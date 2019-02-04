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

	// Add Events
	machine.AddEvent("start", "a")
	machine.AddEvent("to b", "b")
	machine.AddEvent("to d", "d")

	// Add callback for event
	machine.AddCallback("a", func() { fmt.Println("A state: ") })
	machine.AddCallback("b", func() { fmt.Println("B state: ") })

	// Init State
	err = machine.SetStateTransition("a")
	if err != nil {
		fmt.Println(err)
	}

	// Send Event
	err = machine.SendEvent("to b")
	if err != nil {
		fmt.Println(err)
	}
	err = machine.SendEvent("to d")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("State: %s\n", machine.GetCurrentState())
}
