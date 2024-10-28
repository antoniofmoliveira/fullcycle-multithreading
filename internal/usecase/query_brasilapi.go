package usecase

import (
	"context"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

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

func NewQueryBrasilapi(ctx context.Context, cancel context.CancelFunc, cep string) *CepQuery {
	q := &CepQuery{
		Context: ctx,
		Cancel:  cancel,
		Cep:     cep,
		Channel: make(chan dto.Response),
		url:     "https://brasilapi.com.br/api/cep/v1/{{cep}}",
	}
	q.ExtractCepFromBody = BrasilapiExtractCepFromBody
	return q
}
