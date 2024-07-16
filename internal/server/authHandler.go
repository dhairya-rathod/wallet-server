package server

import (
	"encoding/json"
	"net/http"
	"wallet-server/internal/database/models"
	"wallet-server/internal/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	var loginData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		utils.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := s.db.GetUserByEmailOrId(loginData.Email, 0)
	if err != nil {
		utils.SendErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		utils.SendErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	claims := utils.JwtClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	}
	token, err := utils.GenerateToken(claims)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.db.RevokePreviousTokens(user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.db.AddToken(models.Token{
		Token:     token,
		UserId:    user.ID,
		IsRevoked: false,
	}); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["token"] = token
	utils.WriteJsonResponse(w, data, http.StatusOK, "success")
}
