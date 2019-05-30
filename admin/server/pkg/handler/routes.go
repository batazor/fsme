package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github/batazor/fsme/admin/server/pkg/db"
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

	r.Post("/{id}/event/{eventName}", SendEvent)

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get list FSM
	machine, err := db.List()
	if err != nil {
		w.Write([]byte(""))
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
		return;
	}

	w.Write([]byte(string(b)))
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