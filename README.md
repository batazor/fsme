# fsme

Finite State Machine Engine for Go

### Basic Example

From examples/simple.go

```go
package main

import (
	"fmt"
	"github.com/batazor/fsme/fsm"
)

func main() {
	// Create new FSM
	machine, err := fsm.New()

	// Add rule
	machine.AddStateTransitionRules("a", "b", "c")
	machine.AddStateTransitionRules("b", "d", "e")
	machine.AddStateTransitionRules("c", "k")
	machine.AddStateTransitionRules("d", "a")
	machine.AddStateTransitionRules("e", "k")
	machine.AddStateTransitionRules("k")

	// Init State
	err = machine.SetStateTransition("b")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("State: %s\n", machine.GetCurrentState())
}
```

### Extends

1. Start with https://github.com/looplab/fsm
2. Start with https://github.com/theckman/go-fsm

### Callbacks

```
// Add callback for event
machine.AddCallback("start", func(e *fsm.Event) { fmt.Println("a state: " + e.FSM.GetCurrentState()) })
machine.AddCallback("to b", func(e *fsm.Event) { fmt.Println("b state: " + e.FSM.GetCurrentState()) })
```