package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) GetDemoDataHandler(w http.ResponseWriter, r *http.Request) {
	expenses, err := s.db.GetDemoData()
	if err != nil {
		log.Fatal(err)
	}
	jsonResp, err := json.Marshal(expenses)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
