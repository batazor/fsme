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

	// transition is the internal transition functions used either directly
	// or when Transition is called in an asynchronous state transition.
	transition func()

	events map[State]State

	// callbacks maps events and source states to destination statef.
	callbacks map[State]func(*FSM)

	// stateMu guards access to the current state.
	stateMu sync.RWMutex

	// eventMu guards access to Event() and Transition().
	eventMu sync.Mutex
}

// TransitionRuleSet is a set of allowed transitionf. This uses map of struct{}
// to implement a set.
type TransitionRuleSet map[State]struct{}

// Callback is a function type that callbacks should use. Event is the current
// event info as the callback happenf.
type Callback func(*FSM)

// Callbaks is a shorthand for defining the callbacks in FSM
type Callbacks map[State]func(*FSM)

// Event is the info that get passed as a reference in the callbacks.
type Event struct {
	// FSM is a reference to the current FSM.
	FSM *FSM

	// Name is the event name.
	Name State

	// Src is the state before the transition.
	Source State

	// Destination is the state after the transition.
	Destination map[State]struct{}

	// Error is an optional error that can be returned from a callback.
	Error error

	// Args is a optional list of arguments passed to the callback.
	Args []interface{}
}
