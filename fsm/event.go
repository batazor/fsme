package fsm

import (
	"errors"
	"fmt"
)

// AddEvent is a function for adding event
func (f *FSM) AddEvent(name State, dst State) error {
	f.events[name] = dst

	return nil
}

func (f *FSM) SendEvent(event State) error {
	dst := f.events[event]

	//if this is nil we cannot assume any state
	if dst == "" {
		return errors.New(fmt.Sprintf("event %s not register\n", event))
	}

	err := f.SetStateTransition(dst)
	if err != nil {
		return err
	}

	return nil
}
