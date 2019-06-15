package fsm

import "encoding/json"

// Import/Export
// TODO: use DI instead struct Export
type Init interface {
	Export() string
	Import(FSM) error
	// TODO: add work with channels (input/output)
	Watch() chan<- interface{}
	Listen(<-chan interface{})
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

func (f *FSM) Export(exec func(Export, interface{}) interface{}, args interface{}) interface{} {
	str := exec(Export{
		State:       f.state,
		Transitions: f.transitions,
		Events:      f.events,
		Callbacks:   f.callbacks,
	}, args)

	return str
}

func (f *FSM) MarshalJSON() ([]byte, error) {
	return json.Marshal(Export{
		State:       f.state,
		Transitions: f.transitions,
		Events:      f.events,
		Callbacks:   f.callbacks,
	})
}

func (f *FSM) UnmarshalJSON(data []byte) error {
	fsm := Export{
		State:       "",
		Transitions: make(map[State]TransitionRuleSet),
		Callbacks: Callbacks{
			make(map[State]Callback),
			make(map[State]Callback),
		},
		Events: make(map[State]State),
	}

	err := json.Unmarshal(data, &fsm)
	if err != nil {
		return err
	}

	f.state = fsm.State
	f.transitions = fsm.Transitions
	f.events = fsm.Events
	f.callbacks = fsm.Callbacks

	return nil
}
