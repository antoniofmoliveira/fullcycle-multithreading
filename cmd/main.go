package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/report"
	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/usecase"
)

func main() {

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	cep := flag.String("cep", "", "CEP")
	flag.Parse()

	if *cep == "" {
		flag.PrintDefaults()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-termChan
		cancel()
		slog.Info("canceling query")
		os.Exit(0)
	}()

	q0 := usecase.NewQueryBrasilapi(ctx, cancel, *cep)
	q1 := usecase.NewCepQueryViacep(ctx, cancel, *cep)

	go q0.GetCep()
	go q1.GetCep()

	select {
	case <-ctx.Done():
		slog.Info("main: Context deadline exceeded")
	case resBrasilapi := <-q0.Channel:
		report.Report(resBrasilapi.Cep, "Brasilapi")
	case resViaCep := <-q1.Channel:
		report.Report(resViaCep.Cep, "ViaCep")
	}

	time.Sleep(time.Second)

}
