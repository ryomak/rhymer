package util

import (
	"github.com/graphql-go/graphql"
)

var WordType = graphql.NewObject(graphql.ObjectConfig{
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"yomi": &graphql.Field{
			Type: graphql.String,
		},
		"rhymeWord": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var WordField = &graphql.Field{
	Type: WordType,
	Args: graphql.FieldConfigArgument{
		"sentence": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolveWord,
}

var SentenceField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"sentence": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolveSentence,
}

func resolveSentence(p graphql.ResolveParams) (interface{}, error) {
	return p.Args["sentence"], nil
}

func resolveWord(p graphql.ResolveParams) (interface{}, error) {
	arg, _ := p.Args["sentence"].(string)
	return GetNomalRhyme(arg)
}
