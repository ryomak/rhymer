package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryomak/rhymer/server/controller"
	"github.com/ryomak/rhymer/server/util"
	"github.com/urfave/negroni"
)


func Run() {
	config := util.GetConfig()
	db := util.InitDB(config.DBConfig)
	//make server
	rSever := controller.RhymeServer{DB:db}
	//router
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/rhyme", rSever.RhymeHandler)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)
	fmt.Printf("start server :port(%v)", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, n))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	util.JsonErrorResponse(w, http.StatusNotFound, "urlが存在しません")
}
