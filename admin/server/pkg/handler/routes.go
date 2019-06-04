package handler

import (
	"encoding/json"
	"fmt"
	"github.com/batazor/fsme/admin/server/pkg/db"
	"github.com/batazor/fsme/fsm"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
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

	r.Post("/{id}/event/{eventName}", SendEvent)

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get list FSM
	machines, err := db.List()
	if err != nil {
		w.Write([]byte(""))
	}

	response := []interface{}{}
	for _, machine := range machines {
		response = append(response, machine.Export(generateFSM))
	}

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("GetById"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	decoder := json.NewDecoder(r.Body)
	var newFSM Fsm
	err := decoder.Decode(&newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	// Update FSM
	machine, err := db.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Update FSM
	machine.SetStateTransition(newFSM.FSM.State)

	response := machine.Export(generateFSM)

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Delete"))
}

func SendEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventName := chi.URLParam(r, "eventName")
	state := fsm.State(eventName)
	machine, err := db.Get()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = machine.SendEvent(state)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := machine.Export(generateFSM)

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func generateFSM(f fsm.Export) interface{} {
	response := []Fsm{}
	response = append(response, Fsm{
		ID:          "ID_1",
		Description: "Description",
		Title:       "Title",
		FSM:         f,
	})

	return response
}
