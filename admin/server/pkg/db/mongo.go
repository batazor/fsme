package db

import (
	"errors"
	"github.com/batazor/fsme/fsm"
)

var (
	err   error
	state []*fsm.FSM
)

func List() ([]*fsm.FSM, error) {
	// Create new FSM
	newFSM, err := fsm.New()
	if err != nil {
		return nil, err
	}

	// Add rule
	newFSM.AddStateTransitionRules("a", "b", "c")
	newFSM.AddStateTransitionRules("b", "d", "e")
	newFSM.AddStateTransitionRules("c", "k")
	newFSM.AddStateTransitionRules("d", "a")
	newFSM.AddStateTransitionRules("e", "k")
	newFSM.AddStateTransitionRules("k")

	// Add Events
	newFSM.AddEvent("start", "a")
	newFSM.AddEvent("to b", "b")
	newFSM.AddEvent("to d", "d")

	// Init State
	err = newFSM.SetStateTransition("a")
	if err != nil {
		return nil, err
	}

	state := append(state, newFSM)

	return state, nil
}

func Get() (*fsm.FSM, error) {
	if state == nil {
		return nil, errors.New("Not found state")
	}
	return nil, nil
}

func Add(*fsm.FSM, error) {

}
