package util

import (
	"github.com/graphql-go/graphql"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"context"
	"github.com/ryomak/rhymer/server/model"
)
/*******************         global variable      ************************/

func ExecuteQuery(query string, schema graphql.Schema,db *gorm.DB) (*graphql.Result, error) {
	ctx := context.Background()
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:context.WithValue(ctx,model.DBKEY,db),
	})
	if len(r.Errors) > 0 {
		fmt.Println(r.Errors)
		er := ""
		for _, v := range r.Errors {
			er += (v.Message + "\n")
		}
		return nil, errors.New(er)
	}
	return r,nil
}
/*******************         RhymeWord      ************************/

var WordType = graphql.NewObject(graphql.ObjectConfig{
	Name:"word",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"yomi": &graphql.Field{
			Type: graphql.String,
		},
		"rhyme_words": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var WordField = &graphql.Field{
	Type: graphql.NewList(WordType),
	Description:"rhyme word",
	Args: graphql.FieldConfigArgument{
		"sentence": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"convert_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolveWord,
}

func resolveWord(p graphql.ResolveParams) (interface{}, error) {
	db ,ok:= p.Context.Value(model.DBKEY).(*gorm.DB)
	if !ok{
		return nil,errors.New("cannot connect database")
	}
	sentence, ok := p.Args["sentence"].(string)
	if !ok {
		return nil,errors.New("sentensce is empty")
	}
	convertType, ok := p.Args["convert_type"].(string)
	if !ok {
		return nil,errors.New("type is empty")
	}
	switch convertType {
	case "normal":
		return GetNormalRhyme(db, sentence)
	default:
		return GetNormalRhyme(db, sentence)
	}
}
/********************************************************/
