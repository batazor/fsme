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

	// Add callback for event
	machine.AddCallback("start", func(e *fsm.Event) { fmt.Println("a state: " + e.FSM.GetCurrentState()) })
	machine.AddCallback("to b", func(e *fsm.Event) { fmt.Println("b state: " + e.FSM.GetCurrentState()) })

	// Init State
	err = machine.SetStateTransition("b")
	if err != nil {
		fmt.Println(err)
	}

	// Send Event
	err = machine.SendEvent("start")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("State: %s\n", machine.GetCurrentState())
}
