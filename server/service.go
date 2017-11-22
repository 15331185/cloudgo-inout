package service

import (

	"net/http"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"

	"github.com/urfave/negroni"

	"os"

)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options {


		IndentJSON: true,

		Extensions: []string{".gtpl", ".html"},

		Directory: "templates",
	})
	server := negroni.Classic()

	m := mux.NewRouter()

	initRouters(m, formatter)

	server.UseHandler(m)

	return server

}

func initRouters(mx *mux.Router, formatter *render.Render) {

	webRoot := os.Getenv("WEBROOT")

    if len(webRoot) == 0 {

        webRoot = "./assets"

	}

	mx.HandleFunc("/", indexHandlerFunc(formatter)).Methods("GET")

	mx.HandleFunc("/json", apiHandlerFunc(formatter)).Methods("GET")

	mx.HandleFunc("/unknown", unknown)

	mx.HandleFunc("/login", formHandler(formatter)).Methods("POST")
        mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(webRoot))))

}

func indexHandlerFunc(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		formatter.HTML(w, http.StatusOK, "login", "")

	}

}

func apiHandlerFunc(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		formatter.JSON(w, http.StatusOK, struct {

			ID      string `json:"id"`

			Name string `json:"content"`

		} {ID: "8675309", Name: "Hello from Go!"})

	}

}

func unknown(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "555 unknown.", 555)

}

func formHandler(formatter *render.Render) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		id := r.FormValue("id")

		name := r.FormValue("name")

		formatter.HTML(w, http.StatusOK, "detail", struct {

			ID string

			NAME string

		} {ID: id, NAME: name})

	}

}
