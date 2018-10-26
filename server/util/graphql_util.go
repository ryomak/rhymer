package util

import (
	"github.com/graphql-go/graphql"
	"errors"
)

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
	},
	Resolve: resolveWord,
}

func resolveWord(p graphql.ResolveParams) (interface{}, error) {
	arg, ok := p.Args["sentence"].(string)
	if !ok {
		return nil,errors.New("sentensce is empty")
	}
	return GetNormalRhyme(arg)
}
