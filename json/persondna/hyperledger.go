package persondna

// CreatePersonDNA is a document image to be attached to a person's profile
type CreatePersonDNA struct {
	Args []*Request `json:"Args"`
}

// Request provides request type and code for access (if required)
type Request struct {
	Function   string     `json:"type"`
	AccessCode string     `json:"accessCode"`
	Parameter  *PersonDNA `json:"parameter"`
}
