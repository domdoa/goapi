package tools

import (
	"github.com/domdoa/goapi/internal/models"
)

func (d *PostgresDb) GetUserByID(id int) (*models.User, error) {
	query := `SELECT name FROM users WHERE id = $1`

	var name string
	err := d.db.QueryRow(query, id).Scan(&name)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name: name,
	}

	return user, nil

}
