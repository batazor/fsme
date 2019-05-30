package db

import (
	"errors"
	"github.com/batazor/fsme/fsm"
)

var (
	err error
	state *fsm.FSM
)

func List() (*fsm.FSM, error) {
	// Create new FSM
	state, err = fsm.New()
	if err != nil {
		return nil, err
	}

	// Add rule
	state.AddStateTransitionRules("a", "b", "c")
	state.AddStateTransitionRules("b", "d", "e")
	state.AddStateTransitionRules("c", "k")
	state.AddStateTransitionRules("d", "a")
	state.AddStateTransitionRules("e", "k")
	state.AddStateTransitionRules("k")

	// Add Events
	state.AddEvent("start", "a")
	state.AddEvent("to b", "b")
	state.AddEvent("to d", "d")

	// Init State
	err = state.SetStateTransition("a")
	if err != nil {
		return nil, err
	}

	return state, nil
}

func Get() (*fsm.FSM, error) {
	if state == nil {
		return nil, errors.New("Not found state")
	}
	return state, nil
}