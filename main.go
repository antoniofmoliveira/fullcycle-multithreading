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

// getCep performs a GET request on the given url, replacing {{cep}} with the cep argument.
// It returns the response body as a JSON string, or an error if one occurs.
// The context is used to cancel the request if it takes longer than 1 second.
func getCep(ctx context.Context, url string, cep string) (string, error) {
	urlFinal := strings.Replace(url, "{{cep}}", cep, -1)

	req, err := http.NewRequestWithContext(ctx, "GET", urlFinal, nil)
	if err != nil {
		return "", err
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		body, error := io.ReadAll(res.Body)
		if error != nil {
			return "", error
		}
		if strings.Contains(url, "viacep") {
			var v viacep
			error = json.Unmarshal(body, &v)
			if error != nil {
				return "", error
			}
			jsonString, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			return string(jsonString), nil
		} else if strings.Contains(url, "brasilapi") {
			var v brasilapi
			error = json.Unmarshal(body, &v)
			if error != nil {
				return "", error
			}
			jsonString, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			return string(jsonString), nil
		} else {
			return "", err
		}

	}

}

// getCepSub runs getCep in a separate goroutine and sends the result to the given channel.
// If the context is canceled, it prints the error and cancels the context.
// Otherwise, it sends the result to the channel and cancels the context.
func getCepSub(ctx context.Context, cancel context.CancelFunc, ch chan string, url string, cep string) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	default:
		res, err := getCep(ctx, url, cep)
		if err != nil {
			log.Println(err)
		}
		ch <- res
		cancel()
	}
}

// main performs two GET requests to the Brasilapi and ViaCep APIs, for the given cep.
// It runs each request in a separate goroutine and waits for the first one to finish.
// If the context is canceled, it logs the error and closes the channels.
// Otherwise, it logs the result of the first request to finish and closes the channels.
func main() {

	cep := "39408078"

	chViaCep := make(chan string)
	defer close(chViaCep)

	chBrasilapi := make(chan string)
	defer close(chBrasilapi)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go getCepSub(ctx, cancel, chViaCep, urlViacep, cep)
	go getCepSub(ctx, cancel, chBrasilapi, urlBrasilapi, cep)

	select {
	case <-ctx.Done():
		log.Println("Context deadline exceeded")
		log.Println(ctx.Err())
	case resBrasilapi := <-chBrasilapi:
		log.Println("Return from Brasilapi")
		log.Println(resBrasilapi)
	case resViaCep := <-chViaCep:
		log.Println("Return from ViaCep")
		log.Println(resViaCep)
	}
}
