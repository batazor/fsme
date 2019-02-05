package fsm

// Callback is a function type that callbacks should use. Event is the current
// event info as the callback happenf.
type Callback func(*FSM)

// Callbaks is a shorthand for defining the callbacks in FSM
type Callbacks struct {
	enter map[State]Callback
	leave map[State]Callback
}

// AddCallback is a function for adding callback function by event
func (f *FSM) AddCallback(state State, typeCb string, cb func(*FSM)) error {
	switch typeCb {
	case "enter":
		f.callbacks.enter[state] = cb
	case "leave":
		f.callbacks.leave[state] = cb
	}

	return nil
}

// onLeaveCallbacks
func (f *FSM) onLeaveCallbacks(cb func(*FSM)) error {
	if fn, ok := f.callbacks.leave[f.state]; ok {
		fn(f)

		return nil
	}

	return nil
}

// onLeaveCallbacks
func (f *FSM) onEnterCallbacks(cb func(*FSM)) error {
	if fn, ok := f.callbacks.enter[f.state]; ok {
		fn(f)

		return nil
	}

	return nil
}
