package database

import (
	"wallet-server/internal/database/models"
)

func (s *Service) GetUserByToken(token string) (models.User, error) {
	var user models.User

	query :=
		`SELECT user.id, user.name, user.email
		FROM auth_tokens JOIN user ON auth_tokens.user_id = user.id
		WHERE auth_token.token = $1 AND auth_token.is_revoked = false;`
	row, err := s.db.Query(query, token)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return models.User{}, err
	}

	return user, nil
}
