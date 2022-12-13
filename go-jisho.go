package gojisho

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gojp/kana"
)

const (
	api string = "https://jisho.org/api/v1/search/words?keyword="
)

func GetUrl(key string) (string) {
        romaji := kana.KanaToRomaji(key)
        return api+romaji
}
// takes word as key and returns data and error
func Search(key string) (Word, error) {
	var word Word

	url := GetUrl(key) 
	resp, err := http.Get(url)
	if err != nil {
		return word, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return word, err
	}

	json.Unmarshal([]byte(body), &word)

	return word, nil
}

// gets indexes and returns all japanese kanji and writing
func (word Word) TransJapan(index ...int) []Japanese {
	var japanes []Japanese
	for _, val := range index {
                japanes = append(japanes, word.Data[val].Japanese...)
	}
	return japanes
}

func (word Word) EngDefinition(index ...int) []Senses {
	var senses []Senses
	for _, val := range index {
                senses = append(senses, word.Data[val].Senses...)
        }
	return senses
}

func (word Word) Jlpt(index ...int) []string {
        var jlpt []string
        for _, val := range index {
                jlpt = append(jlpt, word.Data[val].Jlpt...)
        }
        return jlpt
}

func (word Word) Status() int {
	return word.Meta.Status
}
