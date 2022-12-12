package gojisho

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	api string = "https://jisho.org/api/v1/search/words?keyword="
)

// takes word as key and returns data and error
func Search(key string) (Word, error) {
	var word Word

	url := func() string {
		return api + key
	}()

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(body), &word)

	return word, nil
}

// gets indexes and returns all japanese kanji and writing
func (word Word) TransJapan(index ...int) []Japanese {
	var data Data
	data = word.Data
	var japanes []Japanese
	for _, val := range index {
		japanes = append(japanes, data[val].Japanese...)
	}
	return japanes
}

func (word Word) EngDefinition(index ...int) []Senses {
	var data Data
	data = word.Data
	var senses []Senses
	for _, val := range index {
		senses = append(senses, data[val].Senses...)
	}
	return senses
}

func (word Word) Status() int {
	return word.Meta.Status
}
