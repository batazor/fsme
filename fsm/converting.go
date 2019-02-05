package fsm

// Save is a function return FSM state machine with all transitions, rules, state
func (f *FSM) Export() FSM {
	return *f
}

// Load is a function for state import
func (f *FSM) Load(sNew *FSM) error {
	f = sNew

	return nil
}
