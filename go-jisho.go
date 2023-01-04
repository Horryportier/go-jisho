package gojisho

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gojp/kana"
)

const (
	api string = "https://jisho.org/api/v1/search/words?keyword="
)

func GetUrl(key string) string {
	return api + key
}

// takes word as key and returns data and error
func Search(key string, isEng bool) (Word, error) {
	var word Word

	if isEng {
		key = kana.KanaToRomaji(key)
	}
	url := GetUrl(key)

	resp, err := http.Get(url)
	if err != nil {
		return word, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return word, err
	}
	defer resp.Body.Close()

	json.Unmarshal([]byte(body), &word)

	return word, nil
}

// gets  indexes and returs respective items
func (word Word) GetEntries(index ...int) []Data {
	var data []Data
	for _, val := range index {
		if val >= word.Len() {
			log.Fatal("index out of range")
		}
		data = append(data, word.Data[val])
	}
	return data
}

// gets indexes and returns all japanese kanji and writing
func (word Word) TransJapan(index ...int) []Japanese {
	var japanes []Japanese
	for _, val := range index {
		if val >= word.Len() {
			log.Fatal("index out of range")
		}
		japanes = append(japanes, word.Data[val].Japanese...)
	}
	return japanes
}

//Gets eng EngDefinition for every item in data
func (word Word) EngDefinition(index ...int) []Senses {
	var senses []Senses
	for _, val := range index {
		if val >= word.Len() {
			log.Fatal("index out of range")
		}
		senses = append(senses, word.Data[val].Senses...)
	}
	return senses
}

//Gets eng Jlpt every item in data
func (word Word) Jlpt(index ...int) []string {
	var jlpt []string
	for _, val := range index {
		if val >= word.Len() {
			log.Fatal("index out of range")
		}
		jlpt = append(jlpt, word.Data[val].Jlpt...)
	}
	return jlpt
}

func (word Word) Status() int {
	return word.Meta.Status
}

func (word Word) Len() int {
	return len(word.Data)
}
