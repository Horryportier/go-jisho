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
func Search(key string) (Data, error) {
	var data Data

	url := func() string {
		return api + key
	}()

        resp, err :=  http.Get(url)
        if err != nil {
                return data, err 
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return data, err
        }

        json.Unmarshal([]byte(body), &data)

	return data, nil
}

// gets indexes and returns all japanes kanji and writing
func (data Data) TransJapan(index ...int) ([]Japanese) {
        var japanes []Japanese
        for _,val := range index{
        japanes = append(japanes, data.Data[val].Japanese...)
        }
        return japanes
}

func (data Data) EngDefinition(index ...int) ([]Senses) {
        var senses []Senses
        for _,val := range index{
                senses = append(senses, data.Data[val].Senses...)
        }
        return senses
}
