package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/antoniofmoliveira/fullcycle-multithreading/internal/dto"
)

func TestGetCepViacep(t *testing.T) {
	type args struct {
		ctx    context.Context
		cancel context.CancelFunc
		cep    string
	}
	Ctx, Cancel := context.WithTimeout(context.Background(), time.Second)
	defer Cancel()
	tests := []struct {
		name string
		args args
		want dto.Response
	}{
		{
			name: "get cep viacep",
			args: args{
				ctx:    Ctx,
				cancel: Cancel,
				cep:    "39408078",
			},
			want: dto.Response{Cep: dto.Cep{
				Cep:          "39408-078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
				Error: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewCepQueryViacep(tt.args.ctx, tt.args.cancel, tt.args.cep)
			go q.GetCep()
			response := <-q.Channel
			if fmt.Sprint(response) != fmt.Sprint(tt.want) {
				t.Errorf("GetCepViacep() = %v, want %v", response, tt.want)
			}
		})
	}
}

func TestGetCepBrasilapi(t *testing.T) {
	type args struct {
		ctx    context.Context
		cancel context.CancelFunc
		cep    string
	}
	Ctx, Cancel := context.WithTimeout(context.Background(), time.Second)
	defer Cancel()
	tests := []struct {
		name string
		args args
		want dto.Response
	}{
		{
			name: "get cep brasilapi",
			args: args{
				ctx:    Ctx,
				cancel: Cancel,
				cep:    "39408078",
			},
			want: dto.Response{Cep: dto.Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
				Error: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueryBrasilapi(tt.args.ctx, tt.args.cancel, tt.args.cep)
			go q.GetCep()
			response := <-q.Channel
			if fmt.Sprint(response) != fmt.Sprint(tt.want) {
				t.Errorf("GetCepBrasilapi() = %v, want %v", response, tt.want)
			}
		})
	}
}
