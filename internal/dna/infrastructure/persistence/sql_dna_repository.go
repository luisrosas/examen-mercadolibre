package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/luisrosas/examen-mercadolibre/internal/dna/domain"

	_ "github.com/go-sql-driver/mysql"
)

type SQLDnaRepository struct {
	db *sql.DB
}

type SQLConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     uint
}

func NewDnaRepository(config SQLConfig) (*SQLDnaRepository, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLDnaRepository{db: db}, nil
}

func (r *SQLDnaRepository) Save(dna domain.Dna) error {
	query := `INSERT INTO dnas (dna, is_mutant) VALUES (?, ?)`
	_, err := r.db.Exec(query, chainToString(dna.Chain), dna.IsMutant())

	return err
}

func (r *SQLDnaRepository) Find(dnaChain []string) (domain.Dna, error) {
	query := `SELECT dna, is_mutant FROM dnas WHERE dna = ? LIMIT 1`

	var dnaRow = DnaModel{}

	row := r.db.QueryRow(query, chainToString(dnaChain))

	err := row.Scan(&dnaRow.Chain, &dnaRow.Mutant)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return domain.Dna{}, domain.ErrNotFound
		default:
			return domain.Dna{}, err
		}
	}

	return domain.Dna{
		Chain:  stringToChain(dnaRow.Chain),
		Mutant: dnaRow.Mutant,
	}, nil
}

func (r *SQLDnaRepository) GetStatistics() (domain.DnaStatistics, error) {
	query := `SELECT
				(SELECT COUNT(*) FROM dnas WHERE is_mutant = true) as mutants,
				(SELECT COUNT(*) FROM dnas WHERE is_mutant = false) as humans`

	var mutants, humans uint
	row := r.db.QueryRow(query)
	err := row.Scan(&mutants, &humans)
	if err != nil {
		return domain.DnaStatistics{}, err
	}

	return domain.DnaStatistics{
		Humans:  humans,
		Mutants: mutants,
	}, nil
}

func chainToString(chain []string) string {
	return strings.Join(chain, "-")
}

func stringToChain(chain string) []string {
	return strings.Split(chain, "-")
}
