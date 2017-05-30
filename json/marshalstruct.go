package json

// Param is a param to be used in JSON types
type Param struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
	Param3 string `json:"param3"`
}

//Person data
type Person struct {
	TaxID string `json:"taxId"`
	Name  string `json:"name"`
	Docs  []Doc  `json:"docs"`
}

//Doc is Person's document data
type Doc struct {
	Type   string `json:"type"`
	Number string `json:"number"`
	Issuer string `json:"issuer"`
}
