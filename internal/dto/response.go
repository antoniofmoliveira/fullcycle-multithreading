package dto

type Response struct {
	Cep   Cep   `json:"cep"`
	Error error `json:"error"`
}

func NewResponse(cep Cep, err error) Response {
	return Response{
		Cep:   cep,
		Error: err,
	}
}