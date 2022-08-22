package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	//mux := http.NewServeMux()
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods(http.MethodGet)

	//r.HandleFunc("/", app.home)
	r.HandleFunc("/", app.home).Methods(http.MethodGet)

	//r.HandleFunc("/snippet/view", app.snippetView)
	r.HandleFunc("/snippet/view/{id}", app.snippetView).Methods(http.MethodGet)

	//r.HandleFunc("/snippet/create", app.snippetCreate)
	r.HandleFunc("/snippet/create", app.snippetCreate).Methods(http.MethodGet)
	r.HandleFunc("/snippet/create", app.snippetCreatePost).Methods(http.MethodPost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(r)
}
