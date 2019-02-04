package fsm

import "fmt"

const (
	callbackNone int = iota
	callbackOnEnter
	callbackOnLeave
)

// AddCallback is a function for adding callback function by event
func (f *FSM) AddCallback(event State, cb func()) error {
	f.callbacks[event] = cb

	return nil
}

// onLeaveCallbacks
func (f *FSM) onLeaveCallbacks(cb func()) error {
	fmt.Println("onLeaveCallbacks", cb)

	if fn, ok := f.callbacks[f.state]; ok {
		fn()

		return nil
	}

	return nil
}

// onLeaveCallbacks
func (f *FSM) onEnterCallbacks(cb func()) error {
	fmt.Println("onEnterCallbacks", cb)

	if fn, ok := f.callbacks[f.state]; ok {
		fn()

		return nil
	}

	return nil
}
