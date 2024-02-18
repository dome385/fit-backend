package dbrepo

import (
	"context"
	"database/sql"
	"fit-backend/internal/models"
	"time"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllÜbungen() ([]*models.Übung, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM uebungen`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var übungen []*models.Übung

	for rows.Next() {
		var übung models.Übung
		err := rows.Scan(
			&übung.ID,
			&übung.Übung,
			&übung.Sätze,
			&übung.Wiederholungen,
			&übung.Gewicht,
		)
		if err != nil {
			return nil, err
		}
		übungen = append(übungen, &übung)
	}

	return übungen, nil
}

func (m *PostgresDBRepo) AddÜbung(übung models.Übung) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into uebungen values ($1, $2, $3, $4, $5) returning id`

	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		übung.Übung,
		übung.Sätze,
		übung.Wiederholungen,
		übung.Gewicht,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}
