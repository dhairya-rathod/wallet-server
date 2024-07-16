package server

import (
	"net/http"
	"wallet-server/internal/utils"
)

func (s *Server) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

	categories, err := s.db.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJsonResponse(w, categories, http.StatusOK, "Categories fetched successfully")
}
