package main

import (
	"fmt"
	"github.com/batazor/fsme/fsm"
)

func genereateDigraph(f fsm.Export, arg interface{}) interface{} {
	str := `digraph StateMachine {
	rankdir=LR
    node[width=1 fixedsize=true shape=circle style=filled fillcolor="darkorchid1" ]
    
`

	for state, transition := range f.Transitions {
		for key, _ := range transition {
			link := fmt.Sprintf("\t%s -> %s\n\r", state, key)
			str = str + link
		}
	}

	str = str + "\n}"

	return str
}

func main() {
	// Create new FSM
	machine, _ := fsm.New()

	// Add rule
	machine.AddStateTransitionRules("a", "b", "c")
	machine.AddStateTransitionRules("b", "d", "e")
	machine.AddStateTransitionRules("c", "k")
	machine.AddStateTransitionRules("d", "a")
	machine.AddStateTransitionRules("e", "k")
	machine.AddStateTransitionRules("k", "k")

	digraph := machine.ExportFunc(genereateDigraph, nil)
	fmt.Println(digraph)
}
