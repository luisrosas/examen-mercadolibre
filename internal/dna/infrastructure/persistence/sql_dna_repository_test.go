package persistence

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"
)

func TestSQLDnaRepository_Save(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		dna domain.Dna
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocker  func(f fields, a args)
		wantErr bool
	}{
		{
			name: "save data",
			fields: fields{
				db: db,
			},
			args: args{
				dna: domain.Dna{
					Chain:  stringToChain("ABC-123"),
					Mutant: true,
				},
			},
			mocker: func(f fields, a args) {
				dbMock.ExpectExec(`INSERT INTO dnas`).WithArgs(
					chainToString(a.dna.Chain),
					a.dna.IsMutant(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields, tt.args)
			r := &SQLDnaRepository{
				db: tt.fields.db,
			}
			if err := r.Save(tt.args.dna); (err != nil) != tt.wantErr {
				t.Errorf("SQLDnaRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := dbMock.ExpectationsWereMet(); err != nil {
				log.Fatalf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSQLDnaRepository_Find(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		db *sql.DB
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
			name:   "find dna",
			fields: fields{db: db},
			args: args{
				dnaChain: stringToChain("ABC-123"),
			},
			mocker: func(f fields, a args) {
				rows := sqlmock.NewRows([]string{"dna", "is_mutant"}).
					AddRow("ABC-123", 1)
				dbMock.ExpectQuery(`SELECT (.+) FROM dnas`).WithArgs(
					chainToString(a.dnaChain),
				).WillReturnRows(rows)
			},
			want: domain.Dna{
				Chain:  stringToChain("ABC-123"),
				Mutant: true,
			},
			wantErr: false,
		},
		{
			name:   "find dna error not found",
			fields: fields{db: db},
			args: args{
				dnaChain: stringToChain("ABC-123"),
			},
			mocker: func(f fields, a args) {
				dbMock.ExpectQuery(`SELECT (.+) FROM dnas`).WithArgs(
					chainToString(a.dnaChain),
				).WillReturnError(sql.ErrNoRows)
			},
			want:    domain.Dna{},
			wantErr: true,
		},
		{
			name:   "find dna error not found",
			fields: fields{db: db},
			args: args{
				dnaChain: stringToChain("ABC-123"),
			},
			mocker: func(f fields, a args) {
				dbMock.ExpectQuery(`SELECT (.+) FROM dnas`).WithArgs(
					chainToString(a.dnaChain),
				).WillReturnError(errors.New("some error"))
			},
			want:    domain.Dna{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields, tt.args)
			r := &SQLDnaRepository{
				db: tt.fields.db,
			}
			got, err := r.Find(tt.args.dnaChain)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDnaRepository.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLDnaRepository.Find() = %v, want %v", got, tt.want)
			}

			if err := dbMock.ExpectationsWereMet(); err != nil {
				log.Fatalf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSQLDnaRepository_GetStatistics(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		mocker  func(f fields)
		want    domain.DnaStatistics
		wantErr bool
	}{
		{
			name:   "get statistics successful",
			fields: fields{db: db},
			mocker: func(f fields) {
				rows := sqlmock.NewRows([]string{"mutants", "humans"}).
					AddRow(6, 2)
				dbMock.ExpectQuery(`SELECT`).WillReturnRows(rows)
			},
			want: domain.DnaStatistics{
				Humans:  2,
				Mutants: 6,
			},
			wantErr: false,
		},
		{
			name:   "get statistics with error",
			fields: fields{db: db},
			mocker: func(f fields) {
				dbMock.ExpectQuery(`SELECT`).WillReturnError(errors.New("some error"))
			},
			want:    domain.DnaStatistics{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			r := &SQLDnaRepository{
				db: tt.fields.db,
			}
			got, err := r.GetStatistics()
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLDnaRepository.GetStatistics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLDnaRepository.GetStatistics() = %v, want %v", got, tt.want)
			}
		})
	}
}
