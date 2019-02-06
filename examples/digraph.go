package main

import (
	"fmt"
	"github.com/batazor/fsme/fsm"
)

func genereateDigraph(f fsm.Export) string {
	str := "digraph G {\n"

	for state, transition := range f.Transitions {
		for key, _ := range transition {
			str = str + "\t" + fmt.Sprintf("%v", state) + "->" + fmt.Sprintf("%v", key) + "\n"
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

	digraph := machine.Export(genereateDigraph)
	fmt.Println(digraph)
}
