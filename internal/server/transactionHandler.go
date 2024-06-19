package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wallet-server/internal/types"
	"wallet-server/internal/utils"

	"github.com/go-chi/chi/v5"
)

func (s *Server) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	user := utils.GetUserFromContext(r.Context())

	transactions, err := s.db.GetTransactions(int(user.UserID))
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp["data"] = transactions
	resp["status"] = http.StatusOK
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}

func (s *Server) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), 400)
		return
	}

	transaction, err := s.db.GetTransaction(id)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp["data"] = transaction
	resp["status"] = http.StatusOK

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}

func (s *Server) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Transaction recorded successfully",
	}

	decoder := json.NewDecoder(r.Body)
	var transaction types.TransactionReq
	err := decoder.Decode(&transaction)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the transaction
	if err := utils.ValidateTransactionReq(transaction); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.AddTransaction(transaction, 1)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
