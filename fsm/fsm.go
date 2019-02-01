package fsm

import (
	"errors"
	"fmt"
)

// New a create a new FSM
func New() (*FSM, error) {
	f := FSM{}

	f.transitions = make(map[State]TransitionRuleSet)
	f.callbacks = make(map[State]Callback)
	f.events = make(map[State]State)

	return &f, nil
}

// AddStateTransitionRules is a function for adding valid state transitions to the machine.
// This allows you to define whicj states any given state can be transitioned to.
func (f *FSM) AddStateTransitionRules(src State, dst ...State) error {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()

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
	f.stateMu.Lock()
	defer f.stateMu.Unlock()

	// if this is nil we cannot assume any state
	if len(f.transitions) == 0 {
		return errors.New("the machine has no states added")
	}

	// if the state is not defined, it's invalid
	if _, ok := f.transitions[toState]; !ok {
		return errors.New(fmt.Sprintf("state %s has not been registered", toState))
	}

	// if the state is nothing, this is probably the inital state
	if f.state == "" {
		// set the state
		f.state = toState
		return nil
	}

	// if we are not permitted to transition to this state...
	if _, ok := f.transitions[f.state][toState]; !ok {
		strError := fmt.Sprintf("transition from state %s to %s is not permitted", f.state, toState)
		return errors.New(strError)
	}

	// set the state
	f.state = toState

	return nil
}

// GetCurrentState returns the current state. If the State returned is
// "", then the machine has not been given an initial state.
func (f *FSM) GetCurrentState() State {
	return f.state
}

// Save is a function return FSM state machine with all transitions, rules, state
func (f *FSM) Save() FSM {
	return *f
}

// Load is a function for state import
func (f *FSM) Load(sNew *FSM) error {
	f = sNew

	return nil
}
