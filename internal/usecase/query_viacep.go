package usecase

import (
	"context"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

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

func NewCepQueryViacep(ctx context.Context, cancel context.CancelFunc, cep string) *CepQuery {
	q := &CepQuery{
		Context: ctx,
		Cancel:  cancel,
		Cep:     cep,
		Channel: make(chan dto.Response),
		url:     "http://viacep.com.br/ws/{{cep}}/json/",
	}
	q.ExtractCepFromBody = ViacepExtractCepFromBody
	return q
}
