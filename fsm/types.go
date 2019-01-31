package fsm

import "sync"

// State is the machine
type State string

// FSM is the state machine that holds the current state.
type FSM struct {
	// current is the state that the FSM is currently in.
	state State

	transitions map[State]TransitionRuleSet
	rules       map[State]map[State]State

	mu sync.Mutex
}

// TransitionRuleSet is a set of allowed transitions. This uses map of struct{}
// to implement a set.
type TransitionRuleSet map[State]struct{}
