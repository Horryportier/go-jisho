package gojisho

type Word struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []Data `json:"data"`
}

type Data struct {
		Slug        string        `json:"slug"`
		IsCommon    bool          `json:"is_common"`
		Tags        []string `json:"tags"`
		Jlpt        []string      `json:"jlpt"`
		Japanese    []Japanese      `json:"japanese"`
		Senses      []Senses        `json:"senses"`
		Attribution Attribution   `json:"attribution"`
}

type Japanese struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

type Senses struct {
	EnglishDefinitions []string      `json:"english_definitions"`
	PartsOfSpeech      []string      `json:"parts_of_speech"`
	Links              []interface{} `json:"links"`
	Tags               []interface{} `json:"tags"`
	Restrictions       []interface{} `json:"restrictions"`
	SeeAlso            []interface{} `json:"see_also"`
	Antonyms           []interface{} `json:"antonyms"`
	Source             []interface{} `json:"source"`
	Info               []interface{} `json:"info"`
}

type Attribution struct {
	Jmdict   bool `json:"jmdict"`
	Jmnedict bool `json:"jmnedict"`
	Dbpedia  bool `json:"dbpedia"`
}
