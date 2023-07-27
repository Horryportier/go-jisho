package gojisho

import (
	"testing"
)


const (
    query = "犬"
)

func TestGetbyKanji(t *testing.T) {
    res, err := Search(query)

    if err != nil {
        t.Errorf("ERROR: failed to search [%s] err: %v", query, err)
    }

    var w WordData 
    err = w.Parse(res)
    if err != nil {
		t.Errorf("ERROR: can't parse data")
    }
}


func TestWordGet(t *testing.T) {
	var w WordData
	
    err := w.Get(query)
    if err != nil {
		t.Errorf("ERROR: couldn't get WordData [%v]", err)
	}
}

func TestWordMethods(t *testing.T) {
    var w WordData 
    err := w.Get(query)
    if err != nil {
		t.Errorf("ERROR: couldn't get WordData [%v]", err)
	}

    d, err := w.First()
	if err != nil {
		t.Errorf("ERROR: couldn't get first item [%v]", err)
	}

    // basic
	if d.Slug != "犬" {
		t.Errorf("ERROR: wrong slug %s should be 犬", d.Slug)
	}
    
    //common
    if w.IsCommon() == false {
        t.Errorf("ERROR: is common should be ture")
    }
    //jap
    jap, err := w.TransJapan([]int{0,2}...)
    if err != nil {
		t.Errorf("ERROR: couldn't get japanes item [%v]", err)
	}
    if jap[0].Reading != "いぬ" {
		t.Errorf("ERROR: japanes Reading should equal `いぬ` not [%s]", jap[0].Reading )
    }


}




