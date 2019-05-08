package fsm

import (
	"encoding/json"
	"fmt"
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

	return r
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := []Fsm{}
	rules := "[]Rules"
	response = append(response, Fsm{
		ID: "ID_1",
		Description: "Description",
		Events: "[]Events",
		Rules: &rules,
		State: "State",
		Title: "Title",
	})
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return;
	}
	fmt.Println(string(b))

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