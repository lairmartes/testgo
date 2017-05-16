package persondna

/*Participante é uma entidade que pode receber os dados fornecidos pelos usuários da rede.
  A lista de participantes é apresentada na tela para o usuário indicar as entidades com
  as quais os dados podem ser compartilhados.*/
type Participante struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Nome    string `json:"nome"`
}
