package domain

import (
	"testing"
)

func TestDna_IsMutant(t *testing.T) {
	type fields struct {
		Chain  []string
		Mutant bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is mutant",
			fields: fields{
				Mutant: true,
			},
			want: true,
		},
		{
			name: "not is mutant",
			fields: fields{
				Mutant: false,
			},
			want: false,
		},
		{
			name: "not is mutant when not set data",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dna{
				Chain:  tt.fields.Chain,
				Mutant: tt.fields.Mutant,
			}
			if got := d.IsMutant(); got != tt.want {
				t.Errorf("Dna.IsMutant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDnaStatistics_Ratio(t *testing.T) {
	type fields struct {
		Humans  uint
		Mutants uint
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "get ratio when count humants is zero",
			fields: fields{
				Mutants: 500,
				Humans:  0,
			},
			want: 500,
		},
		{
			name: "get ratio when count humants not is zero",
			fields: fields{
				Mutants: 500,
				Humans:  250,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DnaStatistics{
				Humans:  tt.fields.Humans,
				Mutants: tt.fields.Mutants,
			}
			if got := d.Ratio(); got != tt.want {
				t.Errorf("DnaStatistics.Ratio() = %v, want %v", got, tt.want)
			}
		})
	}
}
