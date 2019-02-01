package fsm

const (
	callbackNone int = iota
	callbackOnEnter
	callbackOnLeave
)

// AddCallback is a function for adding callback function by event
func (f *FSM) AddCallback(event State, cb Callback) error {
	f.callbacks[event] = cb

	return nil
}
