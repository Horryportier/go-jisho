package gojisho

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	api string = "https://jisho.org/api/v1/search/words?keyword="
)


func getUrl(sentance string) string {
    var parsed string = url.QueryEscape(sentance)
    var url string = fmt.Sprintf("%s%s", api, parsed)
	return url
}

// will give you raw []byte data of the request 
// trows error when bad request 
// you sould only use this method when you wan't raw json data
func Search(sentance string) ([]byte, error) {
	url := getUrl(sentance)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}



// parses byte data into WordData 
// returns error when json.Unmarshal fails
func (word *WordData) Parse(payload []byte) (error) {
	err := json.Unmarshal([]byte(payload), &word)
	if err != nil {
		return  err
	}
	return nil
}


// simpler way to get WordData ex.
//     var w WordData 
//     w.Get("犬")
func (word *WordData) Get(sentace string)  error {
    res,err := Search(sentace)
    if err != nil {
        return err
    }
    err = word.Parse(res)
    if err != nil {
        return err
    }
    return nil
}


// gets  indexes and returs respective items
// returns error when no items in data field 
func (word WordData) GetEntries(index ...int) ([]Data, error) {
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
// returns error when no items in japanese field 
func (word WordData) TransJapan(index ...int) ([]Japanese, error) {
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
// rutuns error when no items in EngDefinition field
func (word WordData) EngDefinition(index ...int) ([]Senses, error) {
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
// rutuns error when no items in Jlpt field
func (word WordData) Jlpt(index ...int) ([]string, error) {
	var data []string
	for _, val := range index {
		if val >= word.Len() {
			return data, errors.New("No items in Jlpt")
		}
		data = append(data, word.Data[val].Jlpt...)
	}
	return data, nil
}

func (word WordData) Status() int {
	return word.Meta.Status
}

// returns leangth of Data instances in WordData 
// usefull when quering with function like Jlpt 
// ex. jlpt := Jlpt(0...Len()) => array of jlpt strings
func (word WordData) Len() int {
	return len(word.Data)
}


// gets First instance of data field from WordData 
func (word *WordData) First() (Data, error) {
	var data Data = Data{}

	d, err := word.GetEntries([]int{1}...)

	if err != nil {
		return data, err
	}
	return d[len(d)-1], nil

}
// get First Transation if theres is none it return empty string
func (word *WordData) FirstTransation() string {
    d, err := word.First()
    if err != nil {
        return ""
    }
    return d.Senses[0].EnglishDefinitions[0]
}

// takse first data field and return IsCommon field 
func (word *WordData) IsCommon() bool {
    d, err := word.First()
    if err != nil  {
        return false
    }
    return d.IsCommon
}

