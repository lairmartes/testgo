package persondna

import "time"

type DocumentoIdentificacao struct {
	Anexo          Anexo     `json:"anexo"`
	Tipo           string    `json:"tipo"`
	Numero         string    `json:"numero"`
	OrgaoExpedidor string    `json:"orgaoExpedidor"`
	DataEmissao    time.Time `json:"dataEmissao"`
}
