package repository

import (
	"database/sql"
	"fit-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllÜbungen() ([]*models.Übung, error)
	AddÜbung(übung models.Übung) (int, error)
}
