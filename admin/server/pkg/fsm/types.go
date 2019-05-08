package fsm

// Fsm fsm
// swagger:model fsm
type Fsm struct {

	// Identity record
	// Read Only: true
	// Format: ObjectId
	ID string `json:"_id,omitempty"`

	// callbacks
	Callbacks string `json:"callbacks,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// events
	Events string `json:"events,omitempty"`

	// rules
	// Required: true
	Rules *string `json:"rules"`

	// state
	State string `json:"state,omitempty"`

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
