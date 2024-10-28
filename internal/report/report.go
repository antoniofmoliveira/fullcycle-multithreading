package report

import (
	"fmt"
	"log/slog"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

func Report(cep dto.Cep, service string) {
	slog.Info("Return from " + service)
	slog.Info(fmt.Sprint(cep))
}