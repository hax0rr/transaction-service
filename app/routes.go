package app

import (
	"github.com/gorilla/mux"
	_ "github.com/hax0rr/transaction-service/docs"
	"github.com/hax0rr/transaction-service/handler"
	"github.com/hax0rr/transaction-service/handler/v1/account"
	"github.com/hax0rr/transaction-service/handler/v1/transaction"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

//	@title			Transaction Service APIs
//	@version		1.0.0
//	@description	This is an API Doc for Transaction Service.
//
//	@host			localhost:8080
//	@BasePath		/
func NewRouter(deps *Dependencies) http.Handler {
	accountHandler := account.NewHandler(deps.AccountService)
	transactionHandler := transaction.NewTransactionHandler(deps.TransactionService, deps.AccountService)

	router := mux.NewRouter()

	router.HandleFunc("/v1/accounts", accountHandler.CreateAccount(deps.Logger)).Methods(http.MethodPost)
	router.HandleFunc("/v1/accounts/{accountID}", accountHandler.GetAccountByID(deps.Logger)).Methods(http.MethodGet)
	router.HandleFunc("/v1/transactions", transactionHandler.CreateTransaction(deps.Logger)).Methods(http.MethodPost)
	router.HandleFunc("/ping", handler.Ping()).Methods(http.MethodGet)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
