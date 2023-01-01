package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("It works!"))
		if err != nil {
			log.Print(err)
		}
	})
	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Print(err)
	}
}
