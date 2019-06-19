package fsm

import (
	"github.com/batazor/fsme/fsm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	FSM         fsm.Export
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Description string
	Name        string
	UI          map[string]UI
}

type PublicFSM struct {
	// current is the state that the FSM is currently in.
	State fsm.State

	// transitions maps events and source states to destination statef.
	Transitions fsm.TransitionRuleSet

	Events map[fsm.State]fsm.State
}

type UI struct {
	X int
	Y int
}

type FSM interface {
	List() ([]*Item, error)
	Get(string) (*Item, error)
	Add(Item) (*primitive.ObjectID, error)
	Update(Item) (*Item, error)
	Delete(string) error
}
