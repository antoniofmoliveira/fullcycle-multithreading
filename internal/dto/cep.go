package dto

import (
	"encoding/json"
	"errors"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/shared"
)

type Cep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func NewCep(cep, state, city, neighborhood, street string) (*Cep, error) {
	c := &Cep{
		Cep:          cep,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Cep) ToJson() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}
	j, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (c *Cep) Validate() error {
	if _, err := shared.ValidateCep(c.Cep); err != nil {
		return err
	}
	if !shared.ValidateStateShort(c.State) {
		return errors.New("state not found")
	}
	if c.City == "" || c.Neighborhood == "" || c.Street == "" {
		return errors.New("city, neighborhood and street must not be empty")
	}
	return nil
}
