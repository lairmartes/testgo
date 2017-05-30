package persondna

// Anexo is a document image to be attached to a person's profile
type Anexo struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	SecureHash string `json:"secureHash"`
	Active     bool   `json:"active"`
}
