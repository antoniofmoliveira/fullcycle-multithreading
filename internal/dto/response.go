package dto

type Response struct {
	Cep   Cep   `json:"cep"`
	Error error `json:"error"`
}

// NewResponse creates a new Response object.
// It takes a Cep object and an error.
// If the error is not nil, it creates a new Response with the error.
// Otherwise, it creates a new Response with the given Cep object.
func NewResponse(cep Cep, err error) Response {
	return Response{
		Cep:   cep,
		Error: err,
	}
}