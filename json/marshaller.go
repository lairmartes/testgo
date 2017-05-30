package json

import (
	"encoding/json"

	"github.com/lairmartes/testgo/json/persondna"
)

// UnmarshalJSONString unmarshals very simple Param
func UnmarshalJSONString(stringInJSONFormat string) (Param, error) {

	var result Param
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

// UnmarshalPersonString unmarshals simple Person
func UnmarshalPersonString(stringInJSONFormat string) (Person, error) {

	var result Person
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

//UnmarshalPersonDNAString unmarshals more complex PersonDNA
func UnmarshalPersonDNAString(stringInJSONFormat string) (persondna.PersonDNA, error) {

	var result persondna.PersonDNA
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}

//UnmarshalCreatePersonDNAArgs unmarshals Hyperledger's create PersonDNA function
func UnmarshalCreatePersonDNAArgs(stringInJSONFormat string) (persondna.CreatePersonDNA, error) {

	var result persondna.CreatePersonDNA
	var err error

	err = json.Unmarshal([]byte(stringInJSONFormat), &result)

	return result, err
}
