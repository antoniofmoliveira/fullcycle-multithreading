package dto

import (
	"reflect"
	"testing"
)

func TestNewCep(t *testing.T) {
	type args struct {
		cep          string
		state        string
		city         string
		neighborhood string
		street       string
	}
	tests := []struct {
		name    string
		args    args
		want    *Cep
		wantErr bool
	}{
		{
			name: "new cep",
			args: args{
				cep:          "39408078",
				state:        "MG",
				city:         "Montes Claros",
				neighborhood: "Ibituruna",
				street:       "Avenida Herlindo Silveira",
			},
			want: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: false,
		},
		{
			name: "new cep with dash",
			args: args{
				cep:          "39408-078",
				state:        "MG",
				city:         "Montes Claros",
				neighborhood: "Ibituruna",
				street:       "Avenida Herlindo Silveira",
			},
			want: &Cep{
				Cep:          "39408-078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: false,
		},
		{
			name: "new cep error",
			args: args{
				cep:          "3940807",
				state:        "MG",
				city:         "Montes Claros",
				neighborhood: "Ibituruna",
				street:       "Avenida Herlindo Silveira",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCep(tt.args.cep, tt.args.state, tt.args.city, tt.args.neighborhood, tt.args.street)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCep_ToJson(t *testing.T) {
	tests := []struct {
		name    string
		c       *Cep
		want    string
		wantErr bool
	}{
		{
			name: "to json",
			c: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			want:    "{\"cep\":\"39408078\",\"state\":\"MG\",\"city\":\"Montes Claros\",\"neighborhood\":\"Ibituruna\",\"street\":\"Avenida Herlindo Silveira\"}",
			wantErr: false,
		},
		{
			name:    "to json error",
			c:       &Cep{},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cep.ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cep.ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCep_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *Cep
		wantErr bool
	}{
		{
			name: "validate cep",
			c: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: false,
		},
		{
			name:    "validate cep error",
			c:       &Cep{},
			wantErr: true,
		},
		{
			name: "validate cep with dash",
			c: &Cep{
				Cep:          "39408-078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: false,
		},

		{
			name: "validate cep error",
			c: &Cep{
				Cep:          "3940807",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: true,
		},

		{
			name: "validate cep with invalid state",
			c: &Cep{
				Cep:          "39408078",
				State:        "MM",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: true,
		},
		{
			name: "validate cep with invalid city",
			c: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: true,
		},

		{
			name: "validate cep with invalid neighborhood",
			c: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "",
				Street:       "Avenida Herlindo Silveira",
			},
			wantErr: true,
		},

		{
			name: "validate cep with invalid street",
			c: &Cep{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Cep.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
