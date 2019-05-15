package handler

import "github.com/batazor/fsme/fsm"

// Fsm fsm
// swagger:model fsm
type Fsm struct {
	FSM fsm.Export

	// Identity record
	// Read Only: true
	// Format: ObjectId
	ID string `json:"_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Event event
// swagger:model event
type Event struct {

	// Identity record
	// Read Only: true
	// Format: ObjectId
	ID string `json:"_id,omitempty"`

	// state
	State string `json:"state,omitempty"`
}
