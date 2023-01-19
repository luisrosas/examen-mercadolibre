package bootstrap

import (
	"os"
	"reflect"
	"testing"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/persistence"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/infrastructure/server"
)

func Test_initDBConfig(t *testing.T) {
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")

	tests := []struct {
		name    string
		setEnv  func()
		want    persistence.SQLConfig
		wantErr bool
	}{
		{
			name:    "no set DB_USER env",
			setEnv:  func() {},
			want:    persistence.SQLConfig{},
			wantErr: true,
		},
		{
			name: "no set DB_PASSWORD env",
			setEnv: func() {
				os.Setenv("DB_USER", "some data")
			},
			want:    persistence.SQLConfig{},
			wantErr: true,
		},
		{
			name: "no set DB_HOST env",
			setEnv: func() {
				os.Setenv("DB_PASSWORD", "some data")
			},
			want:    persistence.SQLConfig{},
			wantErr: true,
		},
		{
			name: "no set DB_PORT env",
			setEnv: func() {
				os.Setenv("DB_HOST", "some data")
			},
			want:    persistence.SQLConfig{},
			wantErr: true,
		},
		{
			name: "no set DB_NAME env",
			setEnv: func() {
				os.Setenv("DB_PORT", "1234")
			},
			want:    persistence.SQLConfig{},
			wantErr: true,
		},
		{
			name: "with all env",
			setEnv: func() {
				os.Setenv("DB_NAME", "some data")
			},
			want: persistence.SQLConfig{
				Host:     "some data",
				User:     "some data",
				Password: "some data",
				Database: "some data",
				Port:     1234,
			},
			wantErr: false,
		},
		{
			name: "env DB_PORT data error",
			setEnv: func() {
				os.Setenv("DB_PORT", "some data")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv()
			got, err := initDBConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("initDBConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initDBConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initServerConfig(t *testing.T) {
	os.Unsetenv("SERVER_PORT")

	tests := []struct {
		name    string
		setEnv  func()
		want    server.Config
		wantErr bool
	}{
		{
			name: "no set SERVER_PORT env",
			setEnv: func() {
			},
			wantErr: true,
		},
		{
			name: "env SERVER_PORT data error",
			setEnv: func() {
				os.Setenv("SERVER_PORT", "some data")
			},
			wantErr: true,
		},
		{
			name: "set SERVER_PORT env",
			setEnv: func() {
				os.Setenv("SERVER_PORT", "8000")
			},
			want:    server.Config{Port: uint(8000)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv()
			got, err := initServerConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("initServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name    string
		setEnv  func()
		wantErr bool
	}{
		{
			name: "error initialize dbconfig",
			setEnv: func() {
				os.Unsetenv("DB_USER")
			},
			wantErr: true,
		},
		{
			name: "error initialize serverconfig",
			setEnv: func() {
				os.Setenv("DB_USER", "some data")
				os.Setenv("DB_PASSWORD", "some data")
				os.Setenv("DB_HOST", "some data")
				os.Setenv("DB_PORT", "1234")
				os.Setenv("DB_NAME", "some data")
				os.Unsetenv("SERVER_PORT")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv()
			_, err := Initialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
