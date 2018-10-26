package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/ryomak/rhymer/server/util"
	"fmt"
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

func executeQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(r.Errors) > 0 {
		fmt.Println(r.Errors)
		er := ""
		for _, v := range r.Errors {
			er += (v.Message + "\n")
		}
		return nil, errors.New(er)
	}
	return r, nil
}

func RhymeHandler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result, err := executeQuery(query, schema)
	if err != nil {
		util.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w = util.WriteJsonHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
