package usecase

import (
	"context"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

// ViacepExtractCepFromBody takes a *CepQuery and a JSON body, attempts to parse it as a dto.Viacep,
// and if successful, converts it to a dto.Cep.
// If the parsing fails, it sends an error to the query's channel and returns an empty dto.Cep and true.
// Otherwise, it returns the converted dto.Cep and false.
func ViacepExtractCepFromBody(c *CepQuery, body []byte) (dto.Cep, bool) {
	cepdto, err := dto.NewViacepFromJson(string(body))
	if err != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, err)
		return dto.Cep{}, true
	}
	cep := dto.Cep{
		Cep:          cepdto.Cep,
		State:        cepdto.Uf,
		City:         cepdto.Localidade,
		Neighborhood: cepdto.Bairro,
		Street:       cepdto.Logradouro,
	}
	return cep, false
}

// NewCepQueryViacep creates a new CepQuery instance configured to use the ViaCEP service.
// It sets up the context, cancel function, cep value, response channel, URL template, and service name.
// It also assigns the ViacepExtractCepFromBody function to handle the extraction of Cep information
// from the response body.
func NewCepQueryViacep(ctx context.Context, cancel context.CancelFunc, cep string) *CepQuery {
	q := &CepQuery{
		Context:     ctx,
		Cancel:      cancel,
		Cep:         cep,
		Channel:     make(chan dto.Response),
		url:         "http://viacep.com.br/ws/{{cep}}/json/",
		ServiceName: "Viacep",
	}
	q.ExtractCepFromBody = ViacepExtractCepFromBody
	return q
}
