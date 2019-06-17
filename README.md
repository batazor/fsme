# fsme

Finite State Machine Engine for Go

[![GoDoc](https://godoc.org/github.com/ryanfaerman/fsm?status.png)](https://godoc.org/github.com/batazor/fsme/fsm)


### Try it

1. [Demo](https://fsme.herokuapp.com/)

-----
![go](./docs/go.png | width=250)
![material-ui](./docs/material-ui.svg | width=250)
![mongo](./docs/mongo-gopher.png | width=250)
![react](./docs/react.png | width=250)
![react-router](./docs/react-router.png | width=250)
![redux](./docs/redux.png | width=250)
-----

### Install

`go get github.com/batazor/fsme/fsm`

### Basic Example

From [examples/simple.go](./examples/simple.go)

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

    // Add Events
    machine.AddEvent("start", "a")
    machine.AddEvent("to b", "b")
    machine.AddEvent("to d", "d")

    // Add callback for state
    machine.AddCallback("a", "enter", func(f *fsm.FSM) { fmt.Println("Enter state: ", f.GetCurrentState()) })
    machine.AddCallback("a", "leave", func(f *fsm.FSM) { fmt.Println("Leave state: ", f.GetCurrentState()) })
    machine.AddCallback("b", "enter", func(f *fsm.FSM) { fmt.Println("Enter state: ", f.GetCurrentState()) })

    // Init State
    err = machine.SetStateTransition("a")
    if err != nil {
        fmt.Println(err)
    }

    // Send Event
    err = machine.SendEvent("to b")
    if err != nil {
        fmt.Println(err)
    }
    err = machine.SendEvent("to d")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("State: %s\n", machine.GetCurrentState())
}
```

### Feature

+ FSM
+ Callback [Entry, Leave]
+ UI Editor/View
+ Export
+ [Examples](./examples)

### Extends

1. Start with https://github.com/looplab/fsm
2. Start with https://github.com/theckman/go-fsm

## UI

![UI](./docs/editor.png)
