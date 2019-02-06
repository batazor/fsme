package fsm

// Import/Export
// TODO: use DI instead struct Export
type Marshalable interface {
	Export() string
	Import() ([]byte, error)
}

// Export this FSM type for export data
type Export struct {
	// current is the state that the FSM is currently in.
	State State

	// transitions maps events and source states to destination statef.
	Transitions map[State]TransitionRuleSet

	Events map[State]State

	// callbacks maps events and source states to destination statef.
	Callbacks Callbacks
}

func (f *FSM) Export(exec func(Export) string) string {
	str := exec(Export{
		State:       f.state,
		Transitions: f.transitions,
		Events:      f.events,
		Callbacks:   f.callbacks,
	})

	return str
}
