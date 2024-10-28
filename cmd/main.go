package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/usecase"
)

func main() {

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

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

	usecase.ExecuteQueries(ctx, cancel, cep)

	time.Sleep(time.Second)

}
