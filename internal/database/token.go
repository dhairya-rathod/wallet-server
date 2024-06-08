package database

import (
	"wallet-server/internal/database/models"
)

// * GetRecordByToken retrieves the demo data from the database.
func (s *Service) GetRecordByToken(tokenString string) (models.Token, error) {
	var t models.Token
	query := `SELECT Id, user_id, is_revoked FROM tokens WHERE token = $1;`

	row, err := s.db.Query(query, tokenString)
	if err != nil {
		return t, err
	}
	defer row.Close()

	if err := row.Scan(&t.ID, &t.UserId, &t.IsRevoked); err != nil {
		return t, err
	}

	return t, nil
}

// * AddToken adds a new token to the database.
func (s *Service) AddToken(token models.Token) error {
	query := `INSERT INTO tokens (token, user_id, is_revoked) VALUES ($1, $2, $3);`

	_, err := s.db.Exec(query, token.Token, token.UserId, token.IsRevoked)
	if err != nil {
		return err
	}
	return nil
}

// * RevokePreviousTokens revokes all previous tokens for the user.
func (s *Service) RevokePreviousTokens(userId int) error {
	query := `UPDATE tokens SET is_revoked = true WHERE user_id = $1;`

	_, err := s.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
