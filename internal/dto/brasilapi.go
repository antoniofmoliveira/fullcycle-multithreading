package dto

import (
	"encoding/json"
	"errors"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/shared"
	"golang.org/x/exp/slices"
)

var services = []string{"viacep", "widenet", "correios", "correios-alt", "open-cep"}

type Brasilapi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilapi(cep, state, city, neighborhood, street, service string) (*Brasilapi, error) {

	b := &Brasilapi{
		Cep:          cep,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		Service:      service,
	}

	err := b.Validate()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewBrasilapiFromJson(jsonString string) (*Brasilapi, error) {
	var b Brasilapi
	err := json.Unmarshal([]byte(jsonString), &b)
	if err != nil {
		return nil, err
	}
	err = b.Validate()
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (b *Brasilapi) Validate() error {
	if _, err := shared.ValidateCepWithoutDash(b.Cep); err != nil {
		return err
	}
	if !shared.ValidateStateShort(b.State) {
		return errors.New("state not found")
	}
	if !slices.Contains(services, b.Service) {
		return errors.New("service not found")
	}
	if b.City == "" || b.Neighborhood == "" || b.Street == "" {
		return errors.New("city, neighborhood and street must not be empty")
	}
	return nil
}
