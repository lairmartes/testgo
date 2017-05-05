package references

import (
	"strings"
	"testing"

	"github.com/lairmartes/testgo/references"
)

const initialValueName = "Mr. Joe Doe"
const initialValueAddress = "Where the streets has no name"
const modifiedValueName = "Mr. Someone"
const modifiedValueAddress = "Where no one lives"
const intReturnedIfStringsAreEquals = 0

func TestReferenceChange(t *testing.T) {

	var myReference references.ReferenceTestType

	myReference.Name = initialValueName
	myReference.Address = initialValueAddress

	t.Log(">>> INFO >>>> Testing initial parameters Name and Address")

	if strings.Compare(myReference.Name, initialValueName) != intReturnedIfStringsAreEquals {
		t.Errorf("Reference attribute is not working. Name should be %s, but is %s.", initialValueName, myReference.Name)
	}

	if strings.Compare(myReference.Address, initialValueAddress) != intReturnedIfStringsAreEquals {
		t.Errorf("Reference attribute is not working for adress.  Address should be %s, but is %s.", initialValueAddress, myReference.Address)
	}

	myReference.ReferenceUpdater(modifiedValueName, modifiedValueAddress)

	if strings.Compare(myReference.Name, modifiedValueName) != intReturnedIfStringsAreEquals {
		t.Errorf("Reference attribute is not working. Name should be %s, but is %s.", modifiedValueName, myReference.Name)
	}

	if strings.Compare(myReference.Address, modifiedValueAddress) != intReturnedIfStringsAreEquals {
		t.Errorf("Reference attribute is not working for adress.  Address should be %s, but is %s.", modifiedValueAddress, myReference.Address)
	}
}
