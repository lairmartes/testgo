package json

import (
	"encoding/json"

	"github.com/lairmartes/testgo/json/persondna"
)

// Unmarshal Param
func UnmarshalJSONString(stringInJSONFormat string) (Param, error) {

	var result Param
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

// Unmarshal Person
func UnmarshalPersonString(stringInJSONFormat string) (Person, error) {

	var result Person
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

//Unmarshal PersonDNA
func UnmarshalPersonDNAString(stringInJSONFormat string) (persondna.PersonDNA, error) {

	var result persondna.PersonDNA
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}
