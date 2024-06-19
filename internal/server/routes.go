package server

import (
	"encoding/json"
	"log"
	"net/http"

	walletMiddleware "wallet-server/internal/middleware"
	"wallet-server/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	// buit-in middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// public routes
	r.Get("/", s.HelloWorldHandler)
	r.Get("/demo-data", s.GetDemoDataHandler)
	r.Get("/categories", s.GetCategoriesHandler)

	// auth routes
	r.Post("/login", s.LoginHandler)

	// private routes
	r.Group(func(r chi.Router) {
		r.Use(walletMiddleware.AuthMiddleware(&s.db))

		// user routes
		r.Route("/user", func(r chi.Router) {
			// r.Patch("/{id}", s.PatchUserHandler)
			// r.Delete("/{id}", s.DeleteUserHandler)
		})

		// transaction routes
		r.Route("/transaction", func(r chi.Router) {
			r.Get("/{id}", s.GetTransactionHandler)
			r.Post("/", s.PostTransactionHandler)
			// r.Patch("/{id}", s.PatchTransactionHandler)
			// r.Delete("/{id}", s.DeleteTransactionHandler)
		})
		r.Get("/transactions", s.GetTransactionsHandler)
	})

	// error handlers
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		utils.SendErrorResponse(w, "Method is not valid", http.StatusMethodNotAllowed)
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"message": "Hello World",
		"status":  200,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
