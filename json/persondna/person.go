package persondna

import (
	"time"
)

//Person is Definição da estrutura que guardará os dados da Pessoa no ledger
type PersonDNA struct {
	ID                      string                   `json:"id"` //verificar se eh realmente necessario ...
	CPF                     string                   `json:"CPF"`
	Nome                    string                   `json:"nome"`
	NomeDoPai               string                   `json:"nomePai"`
	NomeDaMae               string                   `json:"nomeMae"`
	EnderecoCorrespondencia *Endereco                `json:"enderecoCorrespondencia"`
	EnderecoOutro           *Endereco                `json:"enderecoOutro"`
	Telefone1               string                   `json:"telefone1"`
	Telefone2               string                   `json:"telefone2"`
	Nacionalidade           string                   `json:"nacionalidade"`
	DataDeNascimento        time.Time                `json:"dataNascimento"`
	Sexo                    string                   `json:"sexo"`
	EstadoCivil             string                   `json:"estadoCivil"`
	NomeConjuge             string                   `json:"nomeConjuge"`
	Profissao               string                   `json:"profissao"`
	LocalNascimento         string                   `json:"localNascimento"`
	DocumentosIdentificacao []DocumentoIdentificacao `json:"documentosIdentificacao"`
	ComprovanteResidencia   *Anexo                   `json:"comprovanteResidencia"`
	Participantes           []Participante           `json:"participantes"`
}
