package usecase

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

type CepQuery struct {
	Context            context.Context
	Cancel             context.CancelFunc
	Cep                string
	ServiceName        string
	Channel            chan dto.Response
	url                string
	ExtractCepFromBody func(c *CepQuery, body []byte) (dto.Cep, bool)
}

// GetCep executes a GET request on the given cep, using the given context.
// It first waits a random time between 1 and 1500 milliseconds, to simulate
// a real-world scenario.
// If the context is canceled, it prints a message and returns.
// Otherwise, it executes the request and sends the response to the given channel.
func (c *CepQuery) GetCep() {
	time.Sleep(time.Duration(rand.Intn(1500)+1) * time.Millisecond)

	req, shouldReturn := prepareUrl(c)
	if shouldReturn {
		return
	}

	select {
	case <-c.Context.Done():
		slog.Info(c.ServiceName + ": canceled context")
		return
	default:
		executeQuery(req, c)
	}

}

// executeQuery performs an HTTP request using the provided request object and processes the response.
// It sends the result to the CepQuery's channel. If an error occurs during the request, it sends the error to the channel.
// The function handles different HTTP status codes by sending appropriate error messages to the channel.
// In case of a 200 OK status, it processes the response body asynchronously.
func executeQuery(req *http.Request, c *CepQuery) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, err)
		return
	}

	switch res.StatusCode {
	case http.StatusOK:
		go processHttpResponseOk(res, c)
	case http.StatusRequestTimeout:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("time exceeded"))
	case http.StatusNotFound:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("not found"))
	case http.StatusBadRequest:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("cep must have 8 digits"))
	case http.StatusInternalServerError:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("internal server error"))
	case http.StatusServiceUnavailable:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("service unavailable"))
	default:
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("unknown error"))
	}
}

// processHttpResponseOk reads the response body from the given http.Response object,
// checks if the response is a JSON object with an "erro" key set to "true",
// and if it is, sends an error to the channel.
// If the response does not contain the error key, it calls the ExtractCepFromBody method
// to process the response body.
// If the ExtractCepFromBody method returns an error, it sends the error to the channel.
// If the ExtractCepFromBody method returns a Cep object, it sends the object to the channel.
// After sending to the channel, it cancels the context.
func processHttpResponseOk(res *http.Response, c *CepQuery) {
	body, error := io.ReadAll(res.Body)
	if error != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("fail to read the body response: "+error.Error()))
		return
	}

	if strings.Contains(string(body), `"erro": "true"`) {
		c.Channel <- dto.NewResponse(dto.Cep{}, errors.New("not found"))
		return
	}

	cep, shouldReturn := c.ExtractCepFromBody(c, body)
	if shouldReturn {
		return
	}

	c.Channel <- dto.NewResponse(cep, nil)
	c.Cancel()
}

// prepareUrl creates a new HTTP GET request with the given context, using the URL
// stored in the CepQuery object, replacing the "{{cep}}" placeholder with the
// actual cep.
// If the request creation fails, it sends an error to the channel and returns true.
// Otherwise, it returns the created request and false.
func prepareUrl(c *CepQuery) (*http.Request, bool) {
	url := strings.Replace(c.url, "{{cep}}", c.Cep, 1)
	req, err := http.NewRequestWithContext(c.Context, "GET", url, nil)
	if err != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, err)
		return nil, true
	}
	return req, false
}
