package model

import (
	"github.com/jinzhu/gorm"
)

type WordParam struct {
	Name string `json:"name"`
	Yomi string `json:"yomi"`
	RhymeWords []string `json:"rhyme_words"`
}

func GetWordByYomi(db *gorm.DB,yomi string)(*WordParam,error){
	//母音に変換
	word := new(Word)
	if err := db.Preload("RhymeWords").First(&word,"yomi=?",yomi).Error;err != nil{
		if  gorm.IsRecordNotFoundError(err){
			return nil,nil
		}
		return nil,err
	}
	res := WordParam{Name:word.Name,Yomi:word.Yomi}
	for _,v := range word.RhymeWords{
		res.RhymeWords = append(res.RhymeWords,v.Name)
	}
	return &res,nil
}

func AddWord(db *gorm.DB ,Name,Yomi string,rhymes ...string)error{
	word := Word{
		Name:Name,
		Yomi:Yomi,
	}
	for _,v := range rhymes{
		word.RhymeWords = append(word.RhymeWords,RhymeWord{Name:v})
	}
	if err := db.Create(&word).Error;err!=nil{
		return err
	}
	return nil
}
