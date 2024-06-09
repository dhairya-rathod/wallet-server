package server

import (
	"encoding/json"
	"net/http"
	"wallet-server/internal/utils"
)

func (s *Server) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	categories, err := s.db.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp["categories"] = categories
	resp["status"] = http.StatusOK
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
