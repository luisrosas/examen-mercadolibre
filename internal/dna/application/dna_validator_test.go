package application

import (
	"testing"
)

func Test_dnaValidator_IsMutant(t *testing.T) {
	type args struct {
		dnaChain []string
	}
	tests := []struct {
		name string
		d    dnaValidator
		args args
		want bool
	}{
		{
			name: "not mutant chain",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TTGCTT", "CTGTGT", "TTAAGG", "AGACGG", "CCTGTC", "TCTGTT"},
			},
			want: false,
		},
		{
			name: "mutant chain (vertical)",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TAGCTT", "CTGTGT", "TTTTGG", "AGACGG", "CATGTC", "TTCGTT"},
			},
			want: true,
		},
		{
			name: "mutant chain (horizontal)",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TAACTT", "CTATGT", "TTATGG", "AGACGG", "CATGTC", "TTCGTT"},
			},
			want: true,
		},
		{
			name: "mutant chain (oblique)",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TGTCTT", "CTGTGT", "TTAGGG", "AGACGG", "CATGTC", "TTCGTT"},
			},
			want: true,
		},
		{
			name: "mutant chain (oblique invert)",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TATCTT", "CTCTGT", "TCCCGG", "CGACCG", "CATGTC", "TTCGTT"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDnaValidator()
			if got := d.IsMutant(tt.args.dnaChain); got != tt.want {
				t.Errorf("dnaValidator.IsMutant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnaValidator_IsChainValid(t *testing.T) {
	type args struct {
		dnaChain []string
	}
	tests := []struct {
		name string
		d    dnaValidator
		args args
		want bool
	}{
		{
			name: "size chain invalid #1",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TGC", "GAT", "ACG", "TTG"},
			},
			want: false,
		},
		{
			name: "size chain invalid #2",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TGC", "GAT", "ACG"},
			},
			want: false,
		},
		{
			name: "content chain invalid",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"XCCCTT", "CTGTGT", "TTAAGG", "AGACGG", "CCTGTC", "TCTGTT"},
			},
			want: false,
		},
		{
			name: "chain valid",
			d:    dnaValidator{},
			args: args{
				dnaChain: []string{"TCCCTT", "CTGTGT", "TTAAGG", "AGACGG", "CCTGTC", "TCTGTT"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dnaValidator{}
			if got := d.IsChainValid(tt.args.dnaChain); got != tt.want {
				t.Errorf("dnaValidator.IsChainValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
