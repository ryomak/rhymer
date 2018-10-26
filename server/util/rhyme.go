package util

import (
	"errors"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/ikawaha/kagome/tokenizer"
	"strings"
	"github.com/jinzhu/gorm"
	"github.com/ryomak/rhymer/server/model"
)

var t = tokenizer.New()

type AnalysedWord struct {
	Name  string
	Class string
	Yomi  string
}


func Analyse(str string) []AnalysedWord {
	tokens := t.Tokenize(str)
	words := make([]AnalysedWord, 0)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		status := token.Features()
		if status[1] == "数" {
			words = append(words, convertNumberToWord(token.Surface)...)
		} else {
			words = append(words, AnalysedWord{token.Surface, status[0], status[7]})
		}
	}
	return words
}

func FetchRhyme(str string) ([]string, error) {
	doc, err := goquery.NewDocument("https://kujirahand.com/web-tools/Words.php?m=boin-search&opt=comp&key=" + str)
	if err != nil {
		return nil, err
	}
	var rhymeWords []string
	doc.Find("rb").Each(func(_ int, s *goquery.Selection) {
		rhymeWords = append(rhymeWords, s.Text())
	})
	return rhymeWords, nil
}

func GetNormalRhyme(db *gorm.DB, str string) (interface{}, error) {
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
		//DBにあるか確認
		w,err := model.GetWordByYomi(db,v.Yomi)
		if err !=nil {
			er += err.Error()+"\n"
		}
		if w!= nil{
			res=append(res,param{w.Name,w.Yomi,w.RhymeWords})
		}
		//なければ取得
		r, err := FetchRhyme(v.Yomi)
		if err != nil {
			res = append(res, param{v.Name, v.Yomi, nil})
			er += err.Error() + "\n"
		}
		//DB保存
		if err := model.AddWord(db,v.Name,v.Yomi,r...);err != nil{
			er += err.Error()+"\n"
		}
		res = append(res, param{v.Name, v.Yomi, r})
	}
	if er != "" {
		return res, errors.New(er)
	}
	return res, nil
}

/************ util function ***************/
//regexp
var rep = regexp.MustCompile(`[!-/:-~]`)

func IsOnlyJapanese(str string) bool {
	if rep.MatchString(str) {
		return false
	}
	return true
}

func convertNumberToWord(str string)[]AnalysedWord{
	sp := strings.Split(str, "")
	w  := make([]AnalysedWord,0)
	//セパレいと
	for _,s := range sp {
		switch s {
		case "1":
			w = append(w,AnalysedWord{s,"名詞","イチ"})
		case "2":
			w = append(w,AnalysedWord{s,"名詞","ニ"})
		case "3":
			w = append(w,AnalysedWord{s,"名詞","サン"})
		case "4":
			w = append(w,AnalysedWord{s,"名詞","シ"})
		case "5":
			w = append(w,AnalysedWord{s,"名詞","ゴ"})
		case "6":
			w = append(w,AnalysedWord{s,"名詞","ロク"})
		case "7":
			w = append(w,AnalysedWord{s,"名詞","シチ"})
		case "8":
			w = append(w,AnalysedWord{s,"名詞","ハチ"})
		case "9":
			w = append(w,AnalysedWord{s,"名詞","キュウ"})
		case "１":
			w = append(w,AnalysedWord{s,"名詞","イチ"})
		case "２":
			w = append(w,AnalysedWord{s,"名詞","ニ"})
		case "３":
			w = append(w,AnalysedWord{s,"名詞","サン"})
		case "４":
			w = append(w,AnalysedWord{s,"名詞","シ"})
		case "５":
			w = append(w,AnalysedWord{s,"名詞","ゴ"})
		case "６":
			w = append(w,AnalysedWord{s,"名詞","ロク"})
		case "７":
			w = append(w,AnalysedWord{s,"名詞","シチ"})
		case "８":
			w = append(w,AnalysedWord{s,"名詞","ハチ"})
		case "９":
			w = append(w,AnalysedWord{s,"名詞","キュウ"})
		default:
			w = append(w,AnalysedWord{s,"名詞","レイ"})
		}
	}
	return w
}
