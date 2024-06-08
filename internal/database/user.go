package database

import (
	"database/sql"
	"fmt"
	"wallet-server/internal/database/models"
)

func (s *Service) GetUserByEmailOrId(email string, userId int) (models.User, error) {
	var user models.User

	query :=
		`SELECT id, name, email, password
		FROM users WHERE email = $1 OR id = $2;`
	row := s.db.QueryRow(query, email, userId)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, err
	}

	fmt.Printf("User found: %+v\n", user)
	return user, nil
}
