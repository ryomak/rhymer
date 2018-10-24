package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/ryomak/rhymer/server/util"
)

var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(q),
}
var schema, _ = graphql.NewSchema(schemaConfig)

var q graphql.ObjectConfig = graphql.ObjectConfig{
	Fields: &graphql.Fields{
		"sentence": util.SentenceField,
		"word":     util.WordField,
	},
}

func executeQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(r.Errors) > 0 {
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
		util.JsonErrorResponse(w, http.StatusBadRequest, "リクエストエラー")
		return
	}
	w = util.WriteJsonHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
