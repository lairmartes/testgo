package tests

import (
	"strings"
	"testing"

	"github.com/lairmartes/testgo/json"
)

const jsonTestStringValue1 = "value1"
const jsonTestStringValue2 = "value2"
const jsonTestStringValue3 = "value3"
const jsonTestStringParam1 = "\"param1\":\"" + jsonTestStringValue1 + "\""
const jsonTestStringParam2 = "\"param2\":\"" + jsonTestStringValue2 + "\""
const jsonTestStringParam3 = "\"param3\":\"" + jsonTestStringValue3 + "\""
const jsonTestString = "{" + jsonTestStringParam1 + "," + jsonTestStringParam2 + "," + jsonTestStringParam3 + "}"
const intReturnedIfStringsAreEquals = 0

const personTestTaxID = "123.456.789-10"
const personTestName = "João das Neves"
const personTestDocRgType = "RG"
const personTestDocRgNumber = "23.345.567-X"
const personTestDocRgIssuer = "SSP-SP"
const personTestDocTitType = "Titulo Eleitor"
const personTestDocTitNumber = "123.345.567.789.91011"
const personTestDocTitIssuer = "TRE-SP"

func TestJSONConversion(t *testing.T) {

	var err error
	var paramtest json.Param

	paramtest, err = json.UnmarshalJSONString(jsonTestString)

	t.Log(">>> Testing JSON unmarshalling")

	if err != nil {
		t.Errorf("An error ocurred while trying to unmarshal JSON string (%s). Erro: %s", jsonTestString, err)
	} else {

		if strings.Compare(paramtest.Param1, jsonTestStringValue1) != intReturnedIfStringsAreEquals {
			t.Errorf("JSON string %s has been converted in wrong way. Param1 should be %s, but is %s", jsonTestString, jsonTestStringParam1, paramtest.Param1)
		}
		if strings.Compare(paramtest.Param2, jsonTestStringValue2) != intReturnedIfStringsAreEquals {
			t.Errorf("JSON string %s has been converted in wrong way. Param1 should be %s, but is %s", jsonTestString, jsonTestStringParam1, paramtest.Param1)
		}
		if strings.Compare(paramtest.Param3, jsonTestStringValue3) != intReturnedIfStringsAreEquals {
			t.Errorf("JSON string %s has been converted in wrong way. Param1 should be %s, but is %s", jsonTestString, jsonTestStringParam1, paramtest.Param1)
		}
	}
}

func TestJSONPerson(t *testing.T) {

	var err error
	var personTest json.Person
	var personString = buildPersonString()

	personTest, err = json.UnmarshalPersonString(personString)

	if err != nil {
		t.Errorf("An error ocurred while trying to unmarshal JSON string (%s). Erro: %s", personString, err)
	} else {

		showErrorDetailIfAny(t, personString, personTestTaxID, personTest.TaxID)
		showErrorDetailIfAny(t, personString, personTestName, personTest.Name)
		showErrorDetailIfAny(t, personString, personTestDocRgType, personTest.Docs[0].Type)
		showErrorDetailIfAny(t, personString, personTestDocRgNumber, personTest.Docs[0].Number)
		showErrorDetailIfAny(t, personString, personTestDocRgIssuer, personTest.Docs[0].Issuer)
		showErrorDetailIfAny(t, personString, personTestDocTitType, personTest.Docs[1].Type)
		showErrorDetailIfAny(t, personString, personTestDocTitNumber, personTest.Docs[1].Number)
		showErrorDetailIfAny(t, personString, personTestDocTitIssuer, personTest.Docs[1].Issuer)
	}
}

func buildPersonString() string {

	var result = ""

	result = "{" // open person data
	result = result + "\"taxId\":\"" + personTestTaxID + "\",\n"
	result = result + "\"name\":\"" + personTestName + "\",\n"
	result = result + "\"docs\":[{" // open docs
	result = result + "\"type\":\"" + personTestDocRgType + "\",\n"
	result = result + "\"number\":\"" + personTestDocRgNumber + "\",\n"
	result = result + "\"issuer\":\"" + personTestDocRgIssuer + "\"\n"
	result = result + "},\n{" // separate documents
	result = result + "\"type\":\"" + personTestDocTitType + "\",\n"
	result = result + "\"number\":\"" + personTestDocTitNumber + "\",\n"
	result = result + "\"issuer\":\"" + personTestDocTitIssuer + "\"\n"
	result = result + "}]\n" // close docs
	result = result + "}"    // close person data

	return result
}

func showErrorDetailIfAny(t *testing.T, personString, expectedValue, actualValue string) {

	if strings.Compare(expectedValue, actualValue) != intReturnedIfStringsAreEquals {
		t.Errorf("JSON string %s has been converted in wrong way. Tax id shoud be %s, but is %s", personString, expectedValue, actualValue)
	}
}
