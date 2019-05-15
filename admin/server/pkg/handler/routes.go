package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"github.com/batazor/fsme/fsm"
)

// Routes creates a REST router
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))

	r.Get("/", List)
	r.Get("/{id}", GetById)
	r.Post("/", Create)
	r.Patch("/{id}", Update)
	r.Delete("/{id}", Delete)

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Create new FSM
	machine, err := fsm.New()
	if err != nil {
		w.Write([]byte(""))
	}

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

	// Init State
	err = machine.SetStateTransition("a")
	if err != nil {
		fmt.Println(err)
	}

	response := machine.Export(generateFSM)

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return;
	}

	w.Write([]byte(string(b)))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("GetById"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Create"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Update"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Delete"))
}

func generateFSM(f fsm.Export) interface{} {
	response := []Fsm{}
	response = append(response, Fsm{
		ID: "ID_1",
		Description: "Description",
		Title: "Title",
		FSM: f,
	})

	return response
}