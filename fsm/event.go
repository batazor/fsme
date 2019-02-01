package fsm

import (
	"errors"
	"fmt"
)

func (f *FSM) Event(event State, args ...interface{}) error {
	fmt.Println("exec func Event()")

	f.eventMu.Lock()
	defer f.eventMu.Unlock()

	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	// if this is nil we cannot assume any state
	if f.transition == nil {
		return errors.New(fmt.Sprintf("error transition to %s\n", event))
	}

	dst, ok := f.transitions[event]
	if !ok {
		return errors.New("unknown event")
	}

	e := &Event{f, event, f.state, dst, nil, args}

	// Setup the transition, call it later.
	f.transition = func() {
		//f.SetStateTransition(event)

		f.eventOnEnterCallbacks(e)
	}

	if err := f.eventOnLeaveCallbacks(e); err != nil {
		f.transition = nil
		return errors.New(fmt.Sprintf("error eventOnLeaveCallbacks %s", event))
	}

	// Perform the rest of the transition, if not asynchronous.
	//f.stateMu.RUnlock()
	//err = f.doTransition()
	//f.stateMu.RLock()
	//if err != nil {
	//	return errors.New("Internal Error")
	//}

	return nil
}

// AddCallback is a function for adding callback function by event
func (f *FSM) AddEvent(name State, dst State) error {
	f.events[name] = dst

	return nil
}

func (f *FSM) SendEvent(event State) error {
	return nil
}

// Cancel can be called in enter_<EVENT> or leave_<EVENT> to cancel the
// current transition before it happens. It takes an optional error, which will
// overwrite e.Error if set before.
//func (e *Event) Cancel(err ...error) {
//	e.canceled = true
//
//	if len(err) > 0 {
//		e.Error = err[0]
//	}
//}

// Async can be called in leave_<EVENT> to do an asynchronous state transition.
//
// The current state transition will be on hold in the old state until a final
// call to Transition is made. This will comlete the transition and possibly
// call the other callbacks.
//func (e *Event) Async() {
//	e.async = true
//}

// eventOnEnterCallbacks calls the enter_<EVENT> callbacks
func (f *FSM) eventOnEnterCallbacks(e *Event) {
	if fn, ok := f.callbacks[f.state]; ok {
		fn(e)
	}
}

// eventOnLeaveCallbacks calls the leave_<EVENT> callbacks
func (f *FSM) eventOnLeaveCallbacks(e *Event) error {
	if fn, ok := f.callbacks[f.state]; ok {
		fn(e)

		return nil
	}

	return nil
}
