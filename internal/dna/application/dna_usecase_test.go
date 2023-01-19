package application

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/persistence"
)

func TestDnaUseCase_Handle(t *testing.T) {
	var (
		dnaMutant = domain.Dna{
			Chain:  []string{"CCCCTT", "CTGTGT", "TTAAGG", "AGACGG", "CCTGTC", "TCTGTT"},
			Mutant: true,
		}
		dnaHumant = domain.Dna{
			Chain:  []string{"TTGCTT", "CTGTGT", "TTAAGG", "AGACGG", "CCTGTC", "TCTGTT"},
			Mutant: false,
		}
	)
	type fields struct {
		dnaRepository *persistence.MockSQLDnaRepository
		dnaValidator  *MockDnaValidator
	}
	type args struct {
		dnaChain []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(f fields, a args)
		want    domain.Dna
		wantErr bool
	}{
		{
			name: "is dna saved in database",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			args: args{
				dnaChain: []string{"a", "b", "c"},
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(true)
				f.dnaRepository.On("Find", a.dnaChain).Once().Return(dnaMutant, nil)
			},
			want:    dnaMutant,
			wantErr: false,
		},
		{
			name: "is humant dna data",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			args: args{
				dnaChain: dnaHumant.Chain,
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(true)
				f.dnaValidator.On("IsMutant", a.dnaChain).Once().Return(false)
				f.dnaRepository.On("Find", a.dnaChain).Once().Return(domain.Dna{}, domain.ErrNotFound)
				f.dnaRepository.On("Save", dnaHumant).Once().Return(nil)
			},
			want:    dnaHumant,
			wantErr: false,
		},
		{
			name: "is mutant dna data",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			args: args{
				dnaChain: dnaMutant.Chain,
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(true)
				f.dnaValidator.On("IsMutant", a.dnaChain).Once().Return(true)
				f.dnaRepository.On("Find", a.dnaChain).Once().Return(domain.Dna{}, domain.ErrNotFound)
				f.dnaRepository.On("Save", dnaMutant).Once().Return(nil)
			},
			want:    dnaMutant,
			wantErr: false,
		},
		{
			name: "error find in database",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(true)
				f.dnaRepository.On("Find", a.dnaChain).Once().Return(domain.Dna{}, errors.New("some error"))
			},
			want:    domain.Dna{},
			wantErr: true,
		},
		{
			name: "error saving in database",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(true)
				f.dnaValidator.On("IsMutant", a.dnaChain).Once().Return(true)
				f.dnaRepository.On("Find", a.dnaChain).Once().Return(domain.Dna{}, domain.ErrNotFound)
				f.dnaRepository.On("Save", domain.Dna{Mutant: true}).Once().Return(errors.New("some error"))
			},
			want:    domain.Dna{},
			wantErr: true,
		},
		{
			name: "not dna chain valid",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
				dnaValidator:  &MockDnaValidator{},
			},
			mocker: func(f fields, a args) {
				f.dnaValidator.On("IsChainValid", a.dnaChain).Once().Return(false)
			},
			want:    domain.Dna{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields, tt.args)
			d := NewDnaUseCase(tt.fields.dnaRepository, tt.fields.dnaValidator)
			got, err := d.Handle(tt.args.dnaChain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DnaUseCase.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DnaUseCase.Handle() = %v, want %v", got, tt.want)
			}

			tt.fields.dnaRepository.AssertExpectations(t)
			tt.fields.dnaValidator.AssertExpectations(t)
		})
	}
}
