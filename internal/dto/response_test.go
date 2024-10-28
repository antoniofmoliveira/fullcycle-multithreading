package dto

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		cep Cep
		err error
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "new response",
			args: args{
				cep: Cep{
					Cep:          "39408078",
					State:        "MG",
					City:         "Montes Claros",
					Neighborhood: "Ibituruna",
					Street:       "Avenida Herlindo Silveira",
				},
				err: nil,
			},
			want: Response{
				Cep: Cep{
					Cep:          "39408078",
					State:        "MG",
					City:         "Montes Claros",
					Neighborhood: "Ibituruna",
					Street:       "Avenida Herlindo Silveira",
				},
				Error: nil,
			},
		},
		{
			name: "new response error",
			args: args{
				cep: Cep{},
				err: errors.New("error"),
			},
			want: Response{
				Cep:   Cep{},
				Error: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(tt.args.cep, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
