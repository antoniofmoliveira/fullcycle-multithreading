package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const urlBrasilapi = "https://brasilapi.com.br/api/cep/v1/{{cep}}"
const urlViacep = "http://viacep.com.br/ws/{{cep}}/json/"


// getCep performs a GET request to the given url, respecting the given context.
// It returns the response body as a string, or an error if the context is canceled or if the request returns an error.
// It also checks the response status, and returns an error if the status is not 200.
// The error message is based on the status code and the response body.
func getCep(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	select {
	case <-ctx.Done():
		log.Println("canceled context")
		return "", ctx.Err()
	default:
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		switch res.StatusCode {
		case http.StatusOK:
			body, error := io.ReadAll(res.Body)
			if error != nil {
				return "", errors.New("fail to read the body response: " + error.Error())
			}
			if strings.Contains(string(body), "erro") {
				return "", errors.New("not found")
			}
			return string(body), nil
		case http.StatusRequestTimeout:
			return "", errors.New("time exceeded")
		case http.StatusNotFound:
			return "", errors.New("not found")
		case http.StatusBadRequest:
			return "", errors.New("cep must have 8 digits")
		case http.StatusInternalServerError:
			return "", errors.New("internal server error")
		case http.StatusServiceUnavailable:
			return "", errors.New("service unavailable")
		default:
			return "", errors.New("fail to get CEP: " + res.Status)
		}
	}
}

// getCepSub runs getCep in a separate goroutine and sends the result to the given channel.
// If the context is canceled, it prints the error and cancels the context.
// Otherwise, it sends the result to the channel and cancels the context.
func getCepSub(ctx context.Context, cancel context.CancelFunc, ch chan string, url string) {
	select {
	case <-ctx.Done():
		ch <- ctx.Err().Error()
	default:
		res, err := getCep(ctx, url)
		if err != nil {
			ch <- err.Error()
			return
		}
		ch <- res
		cancel()
	}
}

// main performs a GET request on the given cep using both Brasilapi and ViaCEP, and prints the result of the first one to finish.
// It uses a context to cancel the requests if they take longer than 1 second.
// It also captures SIGINT, SIGTERM, and SIGHUP to cancel the execution and print a message.
// It closes the channels used to communicate with the goroutines.
// It logs the error if the context is canceled or if the requests return an error.
// It prints the result of the request if it's successful.
func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <cep>")
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func() {
		<-termChan
        cancel()
		log.Println("canceling query")
		os.Exit(0)
	}()

	cep := os.Args[1]

	chViaCep := make(chan string)
	defer close(chViaCep)

	chBrasilapi := make(chan string)
	defer close(chBrasilapi)

	go getCepSub(ctx, cancel, chViaCep, strings.Replace(urlViacep, "{{cep}}", cep, -1))
	go getCepSub(ctx, cancel, chBrasilapi, strings.Replace(urlBrasilapi, "{{cep}}", cep, -1))

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
