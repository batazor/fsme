package fsm

const (
	callbackNone int = iota
	callbackOnEnter
	callbackOnLeave
)

// AddCallback is a function for adding callback function by event
func (f *FSM) AddCallback(event State, typeCb string, cb func(*FSM)) error {
	// TODO: use typeCb

	f.callbacks[event] = cb

	return nil
}

// onLeaveCallbacks
func (f *FSM) onLeaveCallbacks(cb func(*FSM)) error {
	if fn, ok := f.callbacks[f.state]; ok {
		fn(f)

		return nil
	}

	return nil
}

// onLeaveCallbacks
func (f *FSM) onEnterCallbacks(cb func(*FSM)) error {
	if fn, ok := f.callbacks[f.state]; ok {
		fn(f)

		return nil
	}

	return nil
}
