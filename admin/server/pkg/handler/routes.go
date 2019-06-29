package handler

import (
	"encoding/json"
	"fmt"
	modelFSM "github.com/batazor/fsme/admin/server/pkg/model/fsm"
	"github.com/batazor/fsme/admin/server/pkg/mongo"
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
	r.Post("/", Create)
	r.Patch("/{id}", Update)
	r.Delete("/{id}", Delete)

	r.Post("/{id}/event/{eventName}", SendEvent)

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	machines, err := mongo.Cfg.List()
	if err != nil {
		w.Write([]byte(""))
	}

	b, err := json.Marshal(machines)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	decoder := json.NewDecoder(r.Body)
	var newFSM modelFSM.Item
	err := decoder.Decode(&newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	// Create FSM
	FSM, err := fsm.New()
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error create a new FSM"))
		return
	}

	FSM.Import(newFSM.FSM)

	// Add transition
	// ...code
	// as example
	//FSM.AddStateTransitionRules("a", "b", "c")
	//FSM.AddStateTransitionRules("b", "d", "e")
	// Add event
	//FSM.AddEvent("start", "a")
	//FSM.AddEvent("to b", "b")
	// ...code
	// Add CB
	// ...code
	// Add state
	// ...code

	newFSM.FSM = FSM.Export()

	id, err := mongo.Cfg.Add(newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error create new FSM"))
		return
	}
	newFSM.Id = *id

	b, err := json.Marshal(newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	decoder := json.NewDecoder(r.Body)
	var newFSM modelFSM.Item
	err := decoder.Decode(&newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	newFSM, err = mongo.Cfg.Update(newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error update FSM"))
		return
	}

	b, err := json.Marshal(newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	idFSM := chi.URLParam(r, "id")

	err := mongo.Cfg.Delete(idFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error delete FSM"))
		return
	}

	w.Write([]byte(string("{}")))
}

func SendEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	id := chi.URLParam(r, "id")
	eventName := chi.URLParam(r, "eventName")

	// Get FSM
	newFSM, err := mongo.Cfg.Get(id)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error get FSM"))
		return
	}

	// Create FSM
	FSM, err := fsm.New()
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error create a new FSM"))
		return
	}

	FSM.Import(newFSM.FSM)

	state := fsm.State(eventName)
	err = FSM.SendEvent(state)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	newFSM.FSM = FSM.Export()

	// Update state
	resultFSM, err := mongo.Cfg.Update(*newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error update FSM"))
		return
	}

	b, err := json.Marshal(resultFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	w.Write([]byte(string(b)))
}
