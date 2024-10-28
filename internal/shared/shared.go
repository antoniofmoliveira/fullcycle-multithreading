package shared

import (
	"errors"
	"regexp"

	"golang.org/x/exp/slices"
)

var states_short = []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}

var states_long = []string{"Acre", "Alagoas",
	"Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal",
	"Espirito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
	"Minas Gerais", "Paraí", "Pará", "Paraná", "Pernambuco", "Piaú", "Rio de Janeiro",
	"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
	"São Paulo", "Sergipe", "Tocantins"}

var regioes = []string{"Sul", "Sudeste", "Centro-Oeste", "Norte", "Nordeste"}

func ValidateCep(cep string) (bool, error) {
	var cepRegex = regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !cepRegex.MatchString(cep) {
		return false, errors.New("cep must have 8 digits, optionally with '-'")
	}
	return true, nil
}

func ValidateCepWithDash(cep string) (bool, error) {
	var cepRegex = regexp.MustCompile(`^\d{5}-\d{3}$`)
	if !cepRegex.MatchString(cep) {
		return false, errors.New("cep must have 8 digits, optionally with '-'")
	}
	return true, nil
}

func ValidateCepWithoutDash(cep string) (bool, error) {
	var cepRegex = regexp.MustCompile(`^\d{8}$`)
	if !cepRegex.MatchString(cep) {
		return false, errors.New("cep must have 8 digits, optionally with '-'")
	}
	return true, nil
}

func ValidateStateShort(state string) bool {
	return slices.Contains(states_short, state)
}

func ValidateStateLong(state string) bool {
	return slices.Contains(states_long, state)
}

func ValidateRegiao(regiao string) bool {
	return slices.Contains(regioes, regiao)
}
