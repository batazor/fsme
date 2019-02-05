package fsm

const (
	callbackNone int = iota
	callbackOnEnter
	callbackOnLeave
)

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
