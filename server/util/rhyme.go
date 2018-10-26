package util

import (
	"errors"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/ikawaha/kagome/tokenizer"
	"fmt"
)

var t = tokenizer.New()

type Word struct {
	Name  string
	Class string
	Yomi  string
}


func Analyse(str string) []Word {
	tokens := t.Tokenize(str)
	words := make([]Word, 0)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		status := token.Features()
		if status[1] == "数" {
			words = append(words, Word{token.Surface, status[1], status[7]})
		} else {
			words = append(words, Word{token.Surface, status[0], status[7]})
		}
	}
	return words
}

func FetchRhyme(str string) ([]string, error) {
	fmt.Println(str)
	doc, err := goquery.NewDocument("https://kujirahand.com/web-tools/Words.php?m=boin-search&opt=comp&key=" + str)
	if err != nil {
		return nil, err
	}
	var rhymeWords []string
	doc.Find("rb").Each(func(_ int, s *goquery.Selection) {
		rhymeWords = append(rhymeWords, s.Text())
	})
	fmt.Println(rhymeWords)
	return rhymeWords, nil
}

func GetNormalRhyme(str string) (interface{}, error) {
	type param struct {
		Name       string   `json:"name"`
		Yomi       string   `json:"yomi"`
		RhymeWords []string `json:"rhyme_words"`
	}
	if !IsOnlyJapanese(str) {
		return param{str,"",nil}, errors.New("can use only japanese")
	}
	words := Analyse(str)
	res := make([]param, 0)
	er := ""
	for _, v := range words {
		if v.Class == "数" {
			res = append(res, param{v.Name, v.Yomi, nil})
			continue
		}
		r, err := FetchRhyme(v.Yomi)
		if err != nil {
			res = append(res, param{v.Name, v.Yomi, nil})
			er += err.Error() + "\n"
		}
		res = append(res, param{v.Name, v.Yomi, r})
	}
	if er != "" {
		return res, errors.New(er)
	}
	return res, nil
}

//regexp
var rep = regexp.MustCompile(`[!-/:-~]`)

func IsOnlyJapanese(str string) bool {
	if rep.MatchString(str) {
		return false
	}
	return true
}
