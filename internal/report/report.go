package report

import (
	"log/slog"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

// Report logs a message with the cep and service name.
// It is used to report the result of a query to a CEP service.
func Report(cep dto.Cep, service string) {
	slog.Info("Return from "+service, "cep", cep)
}
