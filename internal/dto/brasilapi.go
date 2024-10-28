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

// NewBrasilapi creates a new Brasilapi instance and validates it.
// It returns an error if the validation fails.
//
// cep must be a valid brazilian cep.
// state must be a valid short state name.
// service must be one of the allowed services.
// city, neighborhood and street must not be empty.
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

// NewBrasilapiFromJson creates a new Brasilapi instance from a JSON string and validates it.
// It returns the created Brasilapi instance or an error if the JSON is invalid or validation fails.
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

// Validate checks the fields of the Brasilapi struct for validity.
// It verifies that the cep is valid, the state is a recognized short state name,
// the service is one of the allowed services, and that the city, neighborhood,
// and street fields are not empty. If any validation fails, it returns an error.
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
