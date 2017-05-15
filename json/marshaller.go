package json

import (
	"encoding/json"
)

func UnmarshalJSONString(stringInJSONFormat string) (Param, error) {

	var result Param
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

func UnmarshalPersonString(stringInJSONFormat string) (Person, error) {

	var result Person
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}
