package gojisho

import (
	"encoding/json"
	"errors"
	"fmt"
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
func Search(key string) ([]byte, error) {
	key = kana.KanaToRomaji(key)
	url := GetUrl(key)

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return body, nil
}

func (word Word) Parse(payload []byte) (Word, error) {
	err := json.Unmarshal([]byte(payload), &word)
	if err != nil {
		fmt.Println("s1 err:", ResolveUnmarshalErr(payload, err))
		return word, err
	}
	return word, nil
}

// gets  indexes and returs respective items
func (word Word) GetEntries(index ...int) ([]Data, error) {
	var data []Data
	for i := range index {
		if i >= word.Len() {
			return data, errors.New("No items in Data")
		}
		data = append(data, word.Data[i])
	}
	return data, nil
}

// gets indexes and returns all japanese kanji and writing
func (word Word) TransJapan(index ...int) ([]Japanese, error) {
	var data []Japanese
	for _, val := range index {
		if val >= word.Len() {
			return data, errors.New("No items in Japanese")
		}
		data = append(data, word.Data[val].Japanese...)
	}
	return data, nil
}

// Gets eng EngDefinition for every item in data
func (word Word) EngDefinition(index ...int) ([]Senses, error) {
	var data []Senses
	for _, val := range index {
		if val >= word.Len() {
			return data, errors.New("No items in Senses")
		}
		data = append(data, word.Data[val].Senses...)
	}
	return data, nil
}

// Gets eng Jlpt every item in data
func (word Word) Jlpt(index ...int) []string {
	var data []string
	for _, val := range index {
		if val >= word.Len() {
			log.Fatal("index out of range")
		}
		data = append(data, word.Data[val].Jlpt...)
	}
	return data
}

func (word Word) Status() int {
	return word.Meta.Status
}

func (word Word) Len() int {
	return len(word.Data)
}

func (word Word) First() (Data, error) {
	var data Data = Data{}

    d, err  := word.GetEntries([]int{1}...)
    if err != nil {
            return data,err
    }
    return  d[len(d)-1] , nil 
}


