package satellite

import (
	"context"
	"database/sql"
	"github.com/BigDwarf/testci/internal/model"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, s model.Satellite) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO satellite (name) VALUES (?)", s.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByName(ctx context.Context, name string) (*model.Satellite, error) {
	var s model.Satellite

	err := r.db.QueryRowContext(ctx, "SELECT * FROM satellite WHERE name = $1", name).Scan(&s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
