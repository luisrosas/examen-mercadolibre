package stats

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	type args struct {
		dnaStatisticsUserCase *application.MockDnaStatisticsUserCase
	}
	tests := []struct {
		name   string
		args   args
		mocker func(a args)
		init   func() int
	}{
		{
			name: "happy path stats",
			args: args{
				dnaStatisticsUserCase: &application.MockDnaStatisticsUserCase{},
			},
			mocker: func(a args) {
				a.dnaStatisticsUserCase.On("Handle", mock.Anything).Once().Return(domain.DnaStatistics{}, nil)
			},
			init: func() int {
				return http.StatusOK
			},
		},
		{
			name: "usecase return error",
			args: args{
				dnaStatisticsUserCase: &application.MockDnaStatisticsUserCase{},
			},
			mocker: func(a args) {
				a.dnaStatisticsUserCase.On("Handle", mock.Anything).Once().Return(
					domain.DnaStatistics{},
					errors.New("some error"),
				)
			},
			init: func() int {
				return http.StatusInternalServerError
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args)
			statusCode := tt.init()

			req := httptest.NewRequest("GET", "http://localhost/stats", &bytes.Buffer{})
			w := httptest.NewRecorder()
			Handle(tt.args.dnaStatisticsUserCase)(w, req)

			assert.True(t, w.Result().StatusCode == statusCode)

			tt.args.dnaStatisticsUserCase.AssertExpectations(t)
		})
	}
}
