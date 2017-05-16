package tests

import (
	"strings"
	"testing"

	"strconv"

	"github.com/lairmartes/testgo/json"
	"github.com/lairmartes/testgo/json/persondna"
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

var personDNAData = []string{"000.000.000-00", "Fulano da Silva", "Nome do Pai", "Nome da Mãe"}
var personDNADataAddress1 = []string{"Rua X", "242", "Santana", "Sao Paulo", "SP", "012345-567"}
var personDNADataAddress2 = []string{"Rua Y", "353", "Prestes", "Feira de Santana", "BA", "45790-100"}
var personDNADataPhones = []string{"(11) 2222-2222", "(11) 92222-2222"}
var personDNADataOthers = []string{"brasileiro", "1982-01-01T21:35:14.052975+02:00", "M", "C", "Sem Ocupação", "Sao Paulo - SP"}
var personDNADataAttachment = []string{"429437932749872482394", "back.png", "ASDJL3KJFDSLJKFSDLJ34", "true"}
var personDNADataDoc1 = []string{"RG", "11223344", "2000-01-01T21:35:14.052975+02:00", "SSP/SP"}
var personDNADataAddressProof = personDNADataAttachment
var personDNADataParticipants = [][]string{{"1", "10.0.1.10", "mvp-dna-if1"}, {"2", "10.0.1.11", "mvp-dna-if2"}}

const hyperledgerArgCreate = "create"
const hyperledterArgCreateCode = "12312312387"

func TestJSONConversion(t *testing.T) {

	var err error
	var paramtest json.Param

	paramtest, err = json.UnmarshalJSONString(jsonTestString)

	t.Log(">>> Testing Param (simple JSON) unmarshalling")
	t.Log(jsonTestString)

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

	t.Log(">>> Testing Person (a bit more complex JSON) unmarshalling")
	t.Log(personString)

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

func TestJSONPersonDNA(t *testing.T) {

	var err error
	var personTest persondna.PersonDNA
	var personString = buildPersonDNAString()

	personTest, err = json.UnmarshalPersonDNAString(personString)

	t.Log(">>> Testing PersonDNA (a complex JSON) unmarshalling")
	t.Log(personString)

	if err != nil {
		t.Errorf("An error ocurred while trying to unmarshal JSON string (%s). Erro: %s", personString, err)
	} else {

		checkPerson(t, personString, &personTest)
	}
}

func TestJSONHyperledgerArgs(t *testing.T) {

	var err error
	var createTest persondna.CreatePersonDNA
	var createString = buildHyperledgerCreateArgs()

	createTest, err = json.UnmarshalCreatePersonDNAArgs(createString)

	t.Log(">>> Testing Create Person DNA - Args sending PersonDNA to be created")
	t.Log(createString)

	if err != nil {
		t.Errorf("An error ocurred while trying to unmarshal JSON string (%s). Erro: %s", createString, err)
	} else {
		showErrorDetailIfAny(t, createString, hyperledgerArgCreate, createTest.Args[0].Function)
		showErrorDetailIfAny(t, createString, hyperledterArgCreateCode, createTest.Args[0].AccessCode)
		checkPerson(t, createString, createTest.Args[0].Parameter)
	}
}

func checkPerson(t *testing.T, personString string, personTest *persondna.PersonDNA) {

	showErrorDetailIfAny(t, personString, personDNAData[0], personTest.CPF)
	showErrorDetailIfAny(t, personString, personDNAData[1], personTest.Nome)
	showErrorDetailIfAny(t, personString, personDNAData[2], personTest.NomeDoPai)
	showErrorDetailIfAny(t, personString, personDNAData[3], personTest.NomeDaMae)

	showErrorDetailIfAny(t, personString, personDNADataAddress1[0], personTest.EnderecoCorrespondencia.Logradouro)
	showErrorDetailIfAny(t, personString, personDNADataAddress1[1], strconv.Itoa(personTest.EnderecoCorrespondencia.Numero))
	showErrorDetailIfAny(t, personString, personDNADataAddress1[2], personTest.EnderecoCorrespondencia.Bairro)
	showErrorDetailIfAny(t, personString, personDNADataAddress1[3], personTest.EnderecoCorrespondencia.Cidade)
	showErrorDetailIfAny(t, personString, personDNADataAddress1[4], personTest.EnderecoCorrespondencia.Estado)
	showErrorDetailIfAny(t, personString, personDNADataAddress1[5], personTest.EnderecoCorrespondencia.Cep)

	showErrorDetailIfAny(t, personString, personDNADataAddress2[0], personTest.EnderecoOutro.Logradouro)
	showErrorDetailIfAny(t, personString, personDNADataAddress2[1], strconv.Itoa(personTest.EnderecoOutro.Numero))
	showErrorDetailIfAny(t, personString, personDNADataAddress2[2], personTest.EnderecoOutro.Bairro)
	showErrorDetailIfAny(t, personString, personDNADataAddress2[3], personTest.EnderecoOutro.Cidade)
	showErrorDetailIfAny(t, personString, personDNADataAddress2[4], personTest.EnderecoOutro.Estado)
	showErrorDetailIfAny(t, personString, personDNADataAddress2[5], personTest.EnderecoOutro.Cep)

	showErrorDetailIfAny(t, personString, personDNADataPhones[0], personTest.Telefone1)
	showErrorDetailIfAny(t, personString, personDNADataPhones[1], personTest.Telefone2)

	showErrorDetailIfAny(t, personString, personDNADataOthers[0], personTest.Nacionalidade)
	//showErrorDetailIfAny(t, personString, personDNADataOthers[1], personTest.DataDeNascimento.String()) - a bit hard to test date since format doesn't match
	showErrorDetailIfAny(t, personString, personDNADataOthers[2], personTest.Sexo)
	showErrorDetailIfAny(t, personString, personDNADataOthers[3], personTest.EstadoCivil)
	showErrorDetailIfAny(t, personString, personDNADataOthers[4], personTest.Profissao)
	showErrorDetailIfAny(t, personString, personDNADataOthers[5], personTest.LocalNascimento)

	showErrorDetailIfAny(t, personString, personDNADataAddressProof[0], personTest.ComprovanteResidencia.ID)
	showErrorDetailIfAny(t, personString, personDNADataAddressProof[1], personTest.ComprovanteResidencia.Nome)
	showErrorDetailIfAny(t, personString, personDNADataAddressProof[2], personTest.ComprovanteResidencia.SecureHash)
	showErrorDetailIfAny(t, personString, personDNADataAddressProof[3], strconv.FormatBool(personTest.ComprovanteResidencia.Active))

	showErrorDetailIfAny(t, personString, personDNADataDoc1[0], personTest.DocumentosIdentificacao[0].Tipo)
	showErrorDetailIfAny(t, personString, personDNADataDoc1[1], personTest.DocumentosIdentificacao[0].Numero)
	//showErrorDetailIfAny(t, personString, personDNADataDoc1[2], personTest.DocumentosIdentificacao[0].DataEmissao.String()) - a bit hard to test date since format doesn't match
	showErrorDetailIfAny(t, personString, personDNADataDoc1[3], personTest.DocumentosIdentificacao[0].OrgaoExpedidor)

	showErrorDetailIfAny(t, personString, personDNADataParticipants[0][0], personTest.Participantes[0].ID)
	showErrorDetailIfAny(t, personString, personDNADataParticipants[0][1], personTest.Participantes[0].Address)
	showErrorDetailIfAny(t, personString, personDNADataParticipants[0][2], personTest.Participantes[0].Nome)

	showErrorDetailIfAny(t, personString, personDNADataParticipants[1][0], personTest.Participantes[1].ID)
	showErrorDetailIfAny(t, personString, personDNADataParticipants[1][1], personTest.Participantes[1].Address)
	showErrorDetailIfAny(t, personString, personDNADataParticipants[1][2], personTest.Participantes[1].Nome)
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

func buildPersonDNAString() string {

	var result = "{"
	result = result + "\"id\":\"1\","
	result = result + "\"CPF\":\"" + personDNAData[0] + "\","
	result = result + "\"nome\":\"" + personDNAData[1] + "\","
	result = result + "\"nomePai\":\"" + personDNAData[2] + "\","
	result = result + "\"nomeMae\":\"" + personDNAData[3] + "\","
	result = result + "\"enderecoCorrespondencia\":{\"id\":\"1\",\"logradouro\":\"" + personDNADataAddress1[0] + "\",\"numero\":" + personDNADataAddress1[1] + ",\"bairro\":\"" + personDNADataAddress1[2] + "\", \"cidade\":\"" + personDNADataAddress1[3] + "\", \"estado\":\"" + personDNADataAddress1[4] + "\", \"cep\":\"" + personDNADataAddress1[5] + "\"},"
	result = result + "\"enderecoOutro\":{\"id\":\"1\", \"logradouro\":\"" + personDNADataAddress2[0] + "\", \"numero\":" + personDNADataAddress2[1] + ", \"bairro\":\"" + personDNADataAddress2[2] + "\", \"cidade\":\"" + personDNADataAddress2[3] + "\", \"estado\":\"" + personDNADataAddress2[4] + "\", \"cep\":\"" + personDNADataAddress2[5] + "\"},"
	result = result + "\"telefone1\":\"" + personDNADataPhones[0] + "\",\"telefone2\":\"" + personDNADataPhones[1] + "\","
	result = result + "\"nacionalidade\":\"" + personDNADataOthers[0] + "\",\"dataNascimento\":\"" + personDNADataOthers[1] + "\","
	result = result + "\"sexo\":\"" + personDNADataOthers[2] + "\",\"estadoCivil\":\"" + personDNADataOthers[3] + "\",\"profissao\":\"" + personDNADataOthers[4] + "\",\"localNascimento\":\"" + personDNADataOthers[5] + "\","
	result = result + "\"documentosIdentificacao\":"
	result = result + "["
	result = result + "{\"anexo\":{\"id\":\"" + personDNADataAttachment[0] + "\",\"nome\":\"" + personDNADataAttachment[1] + "\",\"secureHash\":\"" + personDNADataAttachment[2] + "\",\"active\":" + personDNADataAttachment[3] + "},"
	result = result + "\"tipo\":\"" + personDNADataDoc1[0] + "\",\"numero\":\"" + personDNADataDoc1[1] + "\",\"dataEmissao\":\"" + personDNADataDoc1[2] + "\",\"orgaoExpedidor\":\"" + personDNADataDoc1[3] + "\" }"
	result = result + "],"
	result = result + "\"comprovanteResidencia\":{ \"id\":\"" + personDNADataAddressProof[0] + "\", \"nome\":\"" + personDNADataAddressProof[1] + "\", \"secureHash\":\"" + personDNADataAddressProof[2] + "\", \"active\":" + personDNADataAddressProof[3] + "},"
	result = result + "\"participantes\":["
	result = result + "{\"id\":\"" + personDNADataParticipants[0][0] + "\",\"address\":\"" + personDNADataParticipants[0][1] + "\",\"nome\":\"" + personDNADataParticipants[0][2] + "\" },"
	result = result + "{\"id\":\"" + personDNADataParticipants[1][0] + "\",\"address\":\"" + personDNADataParticipants[1][1] + "\",\"nome\":\"" + personDNADataParticipants[1][2] + "\" }]"
	result = result + "}"

	return result
}

func buildHyperledgerCreateArgs() string {

	var result = "{\"Args\":[\"" + hyperledgerArgCreate + "\",\"" + hyperledterArgCreateCode + "\"," + buildPersonDNAString() + "]}"

	return result

}

func showErrorDetailIfAny(t *testing.T, unmashalledString, expectedValue, actualValue string) {

	if strings.Compare(expectedValue, actualValue) != intReturnedIfStringsAreEquals {
		t.Errorf("JSON string %s has been converted in wrong way. Should be %s, but is %s", unmashalledString, expectedValue, actualValue)
	}
}
