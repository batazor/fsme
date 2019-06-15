package main

import (
	"fmt"
	"github.com/batazor/fsme/fsm"
)

func main() {
	j := `{"State":"a"}`

	f, err := fsm.New()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	f.UnmarshalJSON([]byte(j))
	fmt.Println("STATE:", f.GetCurrentState())
}
