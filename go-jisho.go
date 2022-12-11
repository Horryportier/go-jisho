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
