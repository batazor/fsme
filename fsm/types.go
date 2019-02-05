package fsm

import "sync"

// State is the machine
type State string

// FSM is the state machine that holds the current state.
type FSM struct {
	// current is the state that the FSM is currently in.
	state State

	// transitions maps events and source states to destination statef.
	transitions map[State]TransitionRuleSet

	events map[State]State

	// callbacks maps events and source states to destination statef.
	callbacks Callbacks

	// stateMu guards access to the current state.
	stateMu sync.RWMutex
}

// TransitionRuleSet is a set of allowed transitionf. This uses map of struct{}
// to implement a set.
type TransitionRuleSet map[State]struct{}

// Callback is a function type that callbacks should use. Event is the current
// event info as the callback happenf.
type Callback func(*FSM)

// Callbaks is a shorthand for defining the callbacks in FSM
type Callbacks struct {
	enter map[State]Callback
	leave map[State]Callback
}

type ExportImpl interface {
	export() []byte
}
