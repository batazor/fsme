package handler

import (
	"encoding/json"
	"fmt"
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
	//r.Get("/{id}", GetById)
	r.Post("/", Create)
	r.Patch("/{id}", Update)
	r.Delete("/{id}", Delete)

	r.Post("/{id}/event/{eventName}", SendEvent)

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get list FSM
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

/*
func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("GetById"))
}

*/

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body
	decoder := json.NewDecoder(r.Body)
	var newFSM mongo.FSM
	err := decoder.Decode(&newFSM)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error parse JSON"))
		return
	}

	// Create FSM
	newFSM.FSM, err = fsm.New()
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error create a new FSM"))
		return
	}

	// Add transition
	// ...code
	// as example
	newFSM.FSM.AddStateTransitionRules("a", "b", "c")
	newFSM.FSM.AddStateTransitionRules("b", "d", "e")
	// Add event
	// ...code
	// Add CB
	// ...code
	// Add state
	// ...code

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
	var newFSM mongo.FSM
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
	FSM, err := mongo.Cfg.Get(id)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.Write([]byte("Error get FSM"))
		return
	}

	state := fsm.State(eventName)
	err = FSM.FSM.SendEvent(state)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Update state
	newFSM, err := mongo.Cfg.Update(*FSM)
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
