package gojisho


type Data struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Data []struct {
		Slug     string        `json:"slug"`
		IsCommon bool          `json:"is_common"`
		Tags     []interface{} `json:"tags"`
		Jlpt     []string      `json:"jlpt"`
		Japanese []struct {
			Word    string `json:"word"`
			Reading string `json:"reading"`
		} `json:"japanese"`
		Senses []struct {
			EnglishDefinitions []string      `json:"english_definitions"`
			PartsOfSpeech      []string      `json:"parts_of_speech"`
			Links              []interface{} `json:"links"`
			Tags               []interface{} `json:"tags"`
			Restrictions       []interface{} `json:"restrictions"`
			SeeAlso            []interface{} `json:"see_also"`
			Antonyms           []interface{} `json:"antonyms"`
			Source             []interface{} `json:"source"`
			Info               []interface{} `json:"info"`
		} `json:"senses"`
		Attribution struct {
			Jmdict   bool `json:"jmdict"`
			Jmnedict bool `json:"jmnedict"`
			Dbpedia  bool `json:"dbpedia"`
		} `json:"attribution"`
	} `json:"data"`
}
