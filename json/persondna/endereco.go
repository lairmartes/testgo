package persondna

//Endereco is where someone lives or works
type Endereco struct {
	ID         string `json:"id"`
	Logradouro string `json:"logradouro"`
	Numero     int    `json:"numero"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
	Cep        string `json:"cep"`
}
