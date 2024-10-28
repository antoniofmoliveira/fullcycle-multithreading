package usecase

import (
	"context"
	"log/slog"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/report"
)

// ExecuteQueries starts two goroutines, one for each service, and waits for
// any of them to finish. If the context is canceled, it logs a message and
// exits. If a service returns an error, it logs the error. If a service
// returns a valid response, it reports it.
func ExecuteQueries(ctx context.Context, cancel context.CancelFunc, cep *string) {
	q0 := NewQueryBrasilapi(ctx, cancel, *cep)
	q1 := NewCepQueryViacep(ctx, cancel, *cep)

	go q0.GetCep()
	go q1.GetCep()

	select {
	case <-ctx.Done():
		slog.Info("ExecuteQueries: Context deadline exceeded")
	case r0 := <-q0.Channel:
		if r0.Error != nil {
			slog.Info("main: " + r0.Error.Error())
		}
		report.Report(r0.Cep, q0.ServiceName)
	case r1 := <-q1.Channel:
		if r1.Error != nil {
			slog.Info("main: " + r1.Error.Error())
		}
		report.Report(r1.Cep, q1.ServiceName)
	}
}
