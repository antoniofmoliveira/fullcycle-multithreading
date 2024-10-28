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

// main sets up the logging configuration and parses the command-line argument for the CEP.
// It initializes a context with a timeout of 1 second and sets up signal handling for SIGINT, SIGTERM, and SIGHUP to cancel the ongoing query.
// It executes the queries using the ExecuteQueries function from the usecase package and logs the result.
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
