package fsm

import (
	"errors"
	"fmt"
	"sync"
)

// State is the machine
type State string

// FSM is the state machine that holds the current state.
type FSM struct {
	// Init - set global options
	init Init

	// current is the state that the FSM is currently in.
	state State

	// transitions maps events and source states to destination statef.
	transitions map[State]TransitionRuleSet `bson:"transitions" json:"transitions"`

	events map[State]State `bson:"events" json:"events"`

	// callbacks maps events and source states to destination statef.
	callbacks Callbacks

	// stateMu guards access to the current state.
	stateMu sync.RWMutex
}

// TransitionRuleSet is a set of allowed transitionf. This uses map of struct{}
// to implement a set.
type TransitionRuleSet map[State]struct{}

// Create a new FSM with empty setup
func New() (*FSM, error) {
	f := FSM{
		transitions: make(map[State]TransitionRuleSet),
		callbacks: Callbacks{
			make(map[State]Callback),
			make(map[State]Callback),
		},
		events: make(map[State]State),
	}

	return &f, nil
}

// AddStateTransitionRules is a function for adding valid state transitions to the machine.
// This allows you to define whicj states any given state can be transitioned to.
func (f *FSM) AddStateTransitionRules(src State, dst ...State) error {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()

	// if the transitions map is nil, we need to allocate it
	if f.transitions == nil {
		f.transitions = make(map[State]TransitionRuleSet)
	}

	// if the map fpr the source state does not exist, allocate it
	if f.transitions[src] == nil {
		f.transitions[src] = make(TransitionRuleSet)
	}

	// get a reference to the map we care about
	// avoids doing the map lookup for each iteration
	t := f.transitions[src]

	for _, d := range dst {
		t[d] = struct{}{}
	}

	return nil
}

// SetStateTransition triggers a transition to the toState. This function is also
// used to set the initial stte of machine.
//
// Before you can transition to any state, even for the initial, you stateMust define
// it with AddStateTransitionRules(). If you  are setting the initial state, and that
// state is not define, this will return an ErrInvalidInitialState error.
//
// When transitiong from a state, this function will return an error either
// if the state transition is not allowed, or if the destination state has
// not been defined. In both cases, it's seen as a non-permitted state transition.
func (f *FSM) SetStateTransition(toState State) error {
	// if this is nil we cannot assume any state
	if len(f.transitions) == 0 {
		return errors.New("the machine has no states added")
	}

	// if the state is not defined, it's invalid
	if _, ok := f.transitions[toState]; !ok {
		return errors.New(fmt.Sprintf("state `%s` has not been registered", toState))
	}

	// if the state is nothing, this is probably the inital state
	if f.state == "" {
		f.setState(toState)
		return nil
	}

	// if we are not permitted to transition to this state...
	if _, ok := f.transitions[f.state][toState]; !ok {
		strError := fmt.Sprintf("transition from state %s to %s is not permitted", f.state, toState)
		return errors.New(strError)
	}

	// enter callback
	b := f.callbacks.enter[toState]
	f.onEnterCallbacks(b)

	// leave callback
	a := f.callbacks.leave[f.state]
	f.onLeaveCallbacks(a)

	// set the state
	f.setState(toState)

	return nil
}

// setState is a function for set new state FSM
func (f *FSM) setState(toState State) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()

	f.state = toState
}

// GetCurrentState returns the current state. If the State returned is
// "", then the machine has not been given an initial state.
func (f *FSM) GetCurrentState() State {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()

	return f.state
}
