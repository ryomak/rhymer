package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/ryomak/rhymer/server/util"
)

func init(){
	wordData,err := graphql.NewSchema(schemaConfig)
	if err != nil{
		panic(err)
	}
	schema = wordData
}
var schema graphql.Schema

var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query: rootQuery,
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:"query",
	Fields: graphql.Fields{
		"word":util.WordField,
	},
})



func RhymeHandler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result, err := util.ExecuteQuery(query, schema)
	if err != nil {
		util.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w = util.WriteJsonHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
