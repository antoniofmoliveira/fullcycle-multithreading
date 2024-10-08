package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type brasilapi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type viacep struct {
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

const urlBrasilapi = "https://brasilapi.com.br/api/cep/v1/{{cep}}"
const urlViacep = "http://viacep.com.br/ws/{{cep}}/json/"

func getCepFromViaCep(ctx context.Context, cep string) (*viacep, error) {
	urlFinal := strings.Replace(urlViacep, "{{cep}}", cep, -1)

	req, err := http.NewRequestWithContext(ctx, "GET", urlFinal, nil)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		body, error := io.ReadAll(res.Body)
		if error != nil {
			return nil, error
		}
		var v viacep
		error = json.Unmarshal(body, &v)
		if error != nil {
			return nil, error
		}
		return &v, nil
	}
}

func getCepFromBrasilapi(ctx context.Context, cep string) (*brasilapi, error) {
	urlFinal := strings.Replace(urlBrasilapi, "{{cep}}", cep, -1)

	req, err := http.NewRequestWithContext(ctx, "GET", urlFinal, nil)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		body, error := io.ReadAll(res.Body)
		if error != nil {
			return nil, error
		}
		var v brasilapi
		error = json.Unmarshal(body, &v)
		if error != nil {
			return nil, error
		}
		return &v, nil
	}
}

func main() {

	cep := "39408078"

	chViaCep := make(chan viacep)
	defer close(chViaCep)

	chBrasilapi := make(chan brasilapi)
	defer close(chBrasilapi)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context, ch chan brasilapi) {

		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
		default:
			res, err := getCepFromBrasilapi(ctx, cep)
			if err != nil {
				log.Println(err)
			}
			ch <- *res
			ctx.Done()
		}
	}(ctx, chBrasilapi)

	go func(ctx context.Context, ch chan viacep) {

		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
		default:
			res, err := getCepFromViaCep(ctx, cep)
			if err != nil {
				log.Println(err)
			}
			ch <- *res
			ctx.Done()
		}
	}(ctx, chViaCep)

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case resBrasilapi := <-chBrasilapi:
		log.Println("Return from Brasilapi")
		json.NewEncoder(log.Writer()).Encode(resBrasilapi)
	case resViaCep := <-chViaCep:
		log.Println("Return from ViaCep")
		json.NewEncoder(log.Writer()).Encode(resViaCep)
	}
}
