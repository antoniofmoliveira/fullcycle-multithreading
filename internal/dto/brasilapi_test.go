package dto

import (
	"reflect"
	"testing"
)

func TestNewBrasilapi(t *testing.T) {
	type args struct {
		cep          string
		state        string
		city         string
		neighborhood string
		street       string
		service      string
	}
	tests := []struct {
		name    string
		args    args
		want    *Brasilapi
		wantErr bool
	}{
		{name: "new brasilapi",
			args: args{
				cep:          "39408078",
				state:        "MG",
				city:         "Montes Claros",
				neighborhood: "Ibituruna",
				street:       "Avenida Herlindo Silveira",
				service:      "open-cep",
			},
			want: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: false,
		},

		{name: "new brasilapi error",
			args: args{
				cep:          "3940807",
				state:        "MG",
				city:         "Montes Claros",
				neighborhood: "Ibituruna",
				street:       "Avenida Herlindo Silveira",
				service:      "open-cep",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBrasilapi(tt.args.cep, tt.args.state, tt.args.city, tt.args.neighborhood, tt.args.street, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBrasilapi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrasilapi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBrasilapiFromJson(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		want    *Brasilapi
		wantErr bool
	}{
		{name: "new brasilapi from json",
			args: args{
				jsonString: `{"cep":"39408078","state":"MG","city":"Montes Claros","neighborhood":"Ibituruna","street":"Avenida Herlindo Silveira","service":"open-cep"}`,
			},
			want: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: false,
		},
		{
			name: "new brasilapi from json error",
			args: args{
				jsonString: `{"cep":39408078,"state":"MG","city":"Montes Claros","neighborhood":"Ibituruna","street":"Avenida Herlindo Silveira","service":"open-cep"}`,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "new brasilapi from json error",
			args: args{
				jsonString: `{"cep":"3940807","state":"MG","city":"Montes Claros","neighborhood":"Ibituruna","street":"Avenida Herlindo Silveira","service":"open-cep"}`,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBrasilapiFromJson(tt.args.jsonString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBrasilapiFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrasilapiFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrasilapi_Validate(t *testing.T) {
	tests := []struct {
		name    string
		b       *Brasilapi
		wantErr bool
	}{
		{
			name: "brasilapi validate",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: false,
		},

		{
			name: "brasilapi invalid cep",
			b: &Brasilapi{
				Cep:          "3940807",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: true,
		},
		{
			name: "brasilapi invalid service",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open",
			},
			wantErr: true,
		},
		{
			name: "brasilapi invalid city",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: true,
		},
		{
			name: "brasilapi invalid neighborhood",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: true,
		},
		{
			name: "brasilapi invalid street",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "MG",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "",
				Service:      "open-cep",
			},
			wantErr: true,
		},
		{
			name: "brasilapi invalid state",
			b: &Brasilapi{
				Cep:          "39408078",
				State:        "M",
				City:         "Montes Claros",
				Neighborhood: "Ibituruna",
				Street:       "Avenida Herlindo Silveira",
				Service:      "open-cep",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Brasilapi.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
