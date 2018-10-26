package model

import "github.com/jinzhu/gorm"

type Word struct {
	gorm.Model
	Name string `json:"name"`
	Yomi string `json:"yomi"`
	RhymeWords []RhymeWord `json:"rhyme_words" gorm:"foreignkey:WordID"`
}
type RhymeWord struct {
	gorm.Model
	WordID uint `json:"word_id"`
	Name string `json:"name"`
}