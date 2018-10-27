package my_graphqls

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/ryomak/rhymer/server/util/rhyme"
	"github.com/ryomak/rhymer/server/model"
	"errors"
)
func init(){
	wordData,err := graphql.NewSchema(schemaConfig)
	if err != nil{
		panic(err)
	}
	schema = wordData
}

var schema graphql.Schema

var schemaConfig = graphql.SchemaConfig{
	Query: rootQuery,
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:"query",
	Fields: graphql.Fields{
		"word": WordField,
	},
})

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
		return nil,errors.New("sentence is empty")
	}
	if len([]rune(sentence)) < 2 {
		return nil,errors.New("word number more than 2")
	}
	if len([]rune(sentence)) > 30 {
		return nil,errors.New("word number less than 30")
	}
	convertType, ok := p.Args["convert_type"].(string)
	if !ok {
		return nil,errors.New("type is empty")
	}
	switch convertType {
	case "normal":
		return rhyme.GetNormalRhyme(db, sentence)
	default:
		return rhyme.GetNormalRhyme(db, sentence)
	}
}

func GetWordSchema()graphql.Schema{
	return schema
}