package usecase

import (
	"context"
	"log/slog"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/report"
)

func ExecuteQueries(ctx context.Context, cancel context.CancelFunc, cep *string) {
	q0 := NewQueryBrasilapi(ctx, cancel, *cep)
	q1 := NewCepQueryViacep(ctx, cancel, *cep)

	go q0.GetCep()
	go q1.GetCep()

	select {
	case <-ctx.Done():
		slog.Info("main: Context deadline exceeded")
	case r0 := <-q0.Channel:
		report.Report(r0.Cep, q0.ServiceName)
	case r1 := <-q1.Channel:
		report.Report(r1.Cep, q1.ServiceName)
	}
}
