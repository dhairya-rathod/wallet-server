package database

import (
	"wallet-server/internal/database/models"
)

// GetDemoData retrieves the demo data from the database.
func (s *Service) GetRecordByToken(tokenString string) (models.Token, error) {
	var t models.Token
	query := `SELECT Id, user_id, is_revoked FROM public.token WHERE token = $1;`

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
