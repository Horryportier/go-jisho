package gojisho

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func payload() []byte {
	file, err := os.Open("payload.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var b = make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func TestGetWord(t *testing.T) {
	var w Word
	res, err := Search("inu")
	if err != nil {
		t.Errorf("ERROR: Search failed,%e", err)
	}

	fmt.Printf("p:%v, f:%v", len(res), len(payload())-1)

	if len(res) != len(payload())-1 {
		t.Errorf("ERROR: bad payload data")
	}

	w, err = w.Parse(res)
    if err != nil {
		t.Errorf("ERROR: can't parse data")
    }

	d, err := w.First()
    t.Logf("FIRST: %v", d)
	if err != nil {
		t.Errorf("ERROR: couldn't get first item")
	}
	if d.Slug != "犬" {
		t.Errorf("ERROR: wrong slug %s should be 犬", d.Slug)
	}
    t.Logf("SLUG: %s", d.Slug)
}

//func TestGetFirst(t *testing.T) {
//
//}
