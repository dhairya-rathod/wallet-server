package server

import (
	"encoding/json"
	"net/http"
	"wallet-server/internal/types"
	"wallet-server/internal/utils"
)

func (s *Server) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transactions, err := s.db.GetTransactions(int(user.UserID))
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJsonResponse(w, transactions, http.StatusOK, "success")
}

func (s *Server) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntegerURLParam(r, "id")
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := s.db.GetTransaction(id)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJsonResponse(w, transaction, http.StatusOK, "success")
}

func (s *Server) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {

	user, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var transaction types.TransactionReq
	err = decoder.Decode(&transaction)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the transaction
	if err := utils.ValidateTransactionReq(transaction); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.AddTransaction(transaction, user.UserID)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJsonResponse(w, nil, http.StatusOK, "Transaction recorded successfully")
}

func (s *Server) PatchTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req types.TransactionPatchReq

	id, err := utils.GetIntegerURLParam(r, "id")
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.PatchTransaction(id, req)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJsonResponse(w, nil, http.StatusOK, "Transaction updated successfully")
}

func (s *Server) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIntegerURLParam(r, "id")
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.DeleteTransaction(id)
	if err != nil {
		utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJsonResponse(w, nil, http.StatusOK, "Transaction deleted successfully")
}
