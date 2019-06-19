# fsme

Finite State Machine Engine for Go

[![GoDoc](https://godoc.org/github.com/ryanfaerman/fsm?status.png)](https://godoc.org/github.com/batazor/fsme/fsm)


### Try it

1. [Demo](https://fsme.herokuapp.com/)

-----
<p align="center">
  <img src="./docs/go.png" width="80">
  <img src="./docs/material-ui.svg" width="80">
  <img src="./docs/mongo-gopher.png" width="80">
  <img src="./docs/react.png" width="80">
  <img src="./docs/redux.svg" width="80">
  <img src="./docs/react-router.png" width="80">
</p>

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

## UI

![view](./docs/editor.png)
![view-editor](./docs/view-editor.png)

### Feature

+ FSM
+ Callback [Entry, Leave]
+ UI Editor/View
+ Export
+ [Examples](./examples)

### Refs

- FSM
  - https://github.com/looplab/fsm
  - https://github.com/theckman/go-fsm
