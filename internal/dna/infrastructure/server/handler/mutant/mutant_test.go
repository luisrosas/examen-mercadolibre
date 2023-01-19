package mutant

import (
	"bytes"
	_ "embed"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/application"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:embed golden/dna.json
var requestDna []byte

func TestHandle(t *testing.T) {
	type args struct {
		useCase *application.MockDnaUseCase
	}
	tests := []struct {
		name   string
		args   args
		mocker func(a args)
		init   func() ([]byte, int)
	}{
		{
			name: "happy path mutant",
			args: args{
				useCase: &application.MockDnaUseCase{},
			},
			mocker: func(a args) {
				a.useCase.On("Handle", mock.Anything).Once().Return(domain.Dna{
					Mutant: true,
				}, nil)
			},
			init: func() ([]byte, int) {
				return requestDna, http.StatusOK
			},
		},
		{
			name: "happy path human",
			args: args{
				useCase: &application.MockDnaUseCase{},
			},
			mocker: func(a args) {
				a.useCase.On("Handle", mock.Anything).Once().Return(domain.Dna{
					Mutant: false,
				}, nil)
			},
			init: func() ([]byte, int) {
				return requestDna, http.StatusForbidden
			},
		},
		{
			name: "service response invalid chain",
			args: args{
				useCase: &application.MockDnaUseCase{},
			},
			mocker: func(a args) {
				a.useCase.On("Handle", mock.Anything).Once().Return(domain.Dna{}, application.ErrInvalidChain)
			},
			init: func() ([]byte, int) {
				return requestDna, http.StatusBadRequest
			},
		},
		{
			name: "service response error",
			args: args{
				useCase: &application.MockDnaUseCase{},
			},
			mocker: func(a args) {
				a.useCase.On("Handle", mock.Anything).Once().Return(domain.Dna{}, errors.New("some error"))
			},
			init: func() ([]byte, int) {
				return requestDna, http.StatusInternalServerError
			},
		},
		{
			name: "unmarshal error",
			args: args{
				useCase: &application.MockDnaUseCase{},
			},
			mocker: func(a args) {},
			init: func() ([]byte, int) {
				return []byte(`bad message`), http.StatusInternalServerError
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args)
			body, statusCode := tt.init()

			req := httptest.NewRequest("POST", "http://localhost/mutants", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			Handle(tt.args.useCase)(w, req)

			assert.True(t, w.Result().StatusCode == statusCode)

			tt.args.useCase.AssertExpectations(t)
		})
	}
}
