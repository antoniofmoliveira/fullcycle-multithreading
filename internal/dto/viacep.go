package dto

import (
	"encoding/json"
	"errors"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/shared"
)

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViacep(cep string, logradouro string, complemento string, unidade string, bairro string, localidade string, uf string, estado string, regiao string, ibge string, gia string, ddd string, siafi string) (*Viacep, error) {
	v := &Viacep{
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Unidade:     unidade,
		Bairro:      bairro,
		Localidade:  localidade,
		Uf:          uf,
		Estado:      estado,
		Regiao:      regiao,
		Ibge:        ibge,
		Gia:         gia,
		Ddd:         ddd,
		Siafi:       siafi,
	}
	err := v.Validate()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func NewViacepFromJson(jsonString string) (*Viacep, error) {
	var v Viacep
	err := json.Unmarshal([]byte(jsonString), &v)
	if err != nil {
		return nil, err
	}
	err = v.Validate()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (v *Viacep) Validate() error {
	if _, err := shared.ValidateCepWithDash(v.Cep); err != nil {
		return err
	}
	if !shared.ValidateStateShort(v.Uf) {
		return errors.New("uf not found")
	}
	if !shared.ValidateStateLong(v.Estado) {
		return errors.New("estado not found")
	}
	if !shared.ValidateRegiao(v.Regiao) {
		return errors.New("regiao not found")
	}
	if v.Localidade == "" || v.Bairro == "" || v.Logradouro == "" {
		return errors.New("localidade, bairro and logradouro must not be empty")
	}

	return nil
}
