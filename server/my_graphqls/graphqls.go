package my_graphqls

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/ryomak/rhymer/server/model"
	"context"
	"errors"
)

func ExecuteQuery(query string, schema graphql.Schema,db *gorm.DB) (*graphql.Result, error) {
	ctx := context.Background()
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:context.WithValue(ctx,model.DBKEY,db),
	})
	if len(r.Errors) > 0 {
		er := ""
		for _, v := range r.Errors {
			er += (v.Message + "\n")
		}
		return nil, errors.New(er)
	}
	return r,nil
}