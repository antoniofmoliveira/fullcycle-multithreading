package usecase

import (
	"context"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

// BrasilapiExtractCepFromBody extracts a dto.Cep from the given byte slice, that is assumed to be a JSON
// object from Brasilapi.
// If the body is not a valid JSON, it sends an error to the channel and returns an empty Cep, and true.
// If the body is a valid JSON, but contains a "erro" key set to true, it sends an error to the channel
// and returns an empty Cep, and true.
// If the body is a valid JSON and does not contain a "erro" key set to true, it extracts the cep, state,
// city, neighborhood and street from the JSON, creates a new Cep object and returns it, and false.
func BrasilapiExtractCepFromBody(c *CepQuery, body []byte) (dto.Cep, bool) {
	cepdto, err := dto.NewBrasilapiFromJson(string(body))
	if err != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, err)
		return dto.Cep{}, true
	}
	cep := dto.Cep{
		Cep:          cepdto.Cep,
		State:        cepdto.State,
		City:         cepdto.City,
		Neighborhood: cepdto.Neighborhood,
		Street:       cepdto.Street,
	}
	return cep, false
}

// NewQueryBrasilapi creates a new CepQuery instance configured to use the Brasilapi service.
// It sets up the context, cancel function, cep value, response channel, URL template, and service name.
// It also assigns the BrasilapiExtractCepFromBody function to handle the extraction of Cep information
// from the response body.
func NewQueryBrasilapi(ctx context.Context, cancel context.CancelFunc, cep string) *CepQuery {
	q := &CepQuery{
		Context:     ctx,
		Cancel:      cancel,
		Cep:         cep,
		Channel:     make(chan dto.Response),
		url:         "https://brasilapi.com.br/api/cep/v1/{{cep}}",
		ServiceName: "Brasilapi",
	}
	q.ExtractCepFromBody = BrasilapiExtractCepFromBody
	return q
}
