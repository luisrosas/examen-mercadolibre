package application

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/persistence"
)

func TestDnaStatisticsUserCase_Handle(t *testing.T) {
	dnaStats := domain.DnaStatistics{
		Humans:  300,
		Mutants: 25,
	}

	type fields struct {
		dnaRepository *persistence.MockSQLDnaRepository
	}
	tests := []struct {
		name    string
		fields  fields
		mocker  func(f fields)
		want    domain.DnaStatistics
		wantErr bool
	}{
		{
			name: "get statistics successful",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
			},
			mocker: func(f fields) {

				f.dnaRepository.On("GetStatistics").Once().Return(dnaStats, nil)
			},
			want:    dnaStats,
			wantErr: false,
		},
		{
			name: "get statistics with error",
			fields: fields{
				dnaRepository: &persistence.MockSQLDnaRepository{},
			},
			mocker: func(f fields) {

				f.dnaRepository.On("GetStatistics").Once().Return(domain.DnaStatistics{}, errors.New(""))
			},
			want:    domain.DnaStatistics{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			d := NewDnaStatsUserCase(tt.fields.dnaRepository)
			got, err := d.Handle()
			if (err != nil) != tt.wantErr {
				t.Errorf("DnaStatisticsUserCase.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DnaStatisticsUserCase.Handle() = %v, want %v", got, tt.want)
			}

			tt.fields.dnaRepository.AssertExpectations(t)
		})
	}
}
