package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryomak/rhymer/server/controller"
	"github.com/ryomak/rhymer/server/util"
	"github.com/urfave/negroni"
)

func Run() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/rhyme", controller.RhymeHandler)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)
	http.ListenAndServe(":8080", n)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	util.JsonErrorResponse(w, http.StatusNotFound, "urlが存在しません")
}
