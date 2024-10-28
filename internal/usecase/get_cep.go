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

func (c *CepQuery) GetCep() {
	randomInt := rand.Intn(1500) + 1
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
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

func prepareUrl(c *CepQuery) (*http.Request, bool) {
	url := strings.Replace(c.url, "{{cep}}", c.Cep, 1)
	req, err := http.NewRequestWithContext(c.Context, "GET", url, nil)
	if err != nil {
		c.Channel <- dto.NewResponse(dto.Cep{}, err)
		return nil, true
	}
	return req, false
}
