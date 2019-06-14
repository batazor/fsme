package mongo

import (
	"github.com/batazor/fsme/fsm"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Client *mongo.Client
}

type FSM struct {
	FSM         *fsm.FSM
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Description string
	Name        string
}

type PublicFSM struct {
	// current is the state that the FSM is currently in.
	State fsm.State

	// transitions maps events and source states to destination statef.
	Transitions fsm.TransitionRuleSet

	Events map[fsm.State]fsm.State
}
