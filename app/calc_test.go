package app

import (
	"reflect"
	"testing"
)

func TestCalc_Sum(t *testing.T) {
	type fields struct {
		A float64
		B float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "teste 1",
			fields: fields{
				A: 200.2,
				B: 200.4,
			},
			want: 400.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calc{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := c.Sum(); got != tt.want {
				t.Errorf("Calc.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCalc(t *testing.T) {
	tests := []struct {
		name string
		want Calc
	}{
		{
			name: "teste 2",
			want: Calc{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCalc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
