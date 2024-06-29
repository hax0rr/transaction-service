package transaction

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"net/http"
	"transaction-service/handler"
	"transaction-service/internal/entity"
	"transaction-service/internal/service/account"
	"transaction-service/internal/service/transaction"
)

//
//type Handler struct {
//	service transaction.IService
//}
//
//func NewHandler(service Service) *Handler {
//	return &Handler{
//		service: service,
//	}
//}

type Handler struct {
	TxnSvc     transaction.IService
	AccountSvc account.IService
}

func NewTransactionHandler(service transaction.IService, svc account.IService) *Handler {
	return &Handler{TxnSvc: service, AccountSvc: svc}
}

// CreateTransaction
//
//	@Summary	Create transaction
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		transaction		body		CreateTransactionRequest	true	"Transaction object"
//	@Success	200				{object}	handler.Response
//	@Failure		404 {object} 	handler.Response
//	@Failure		400 {object} 	handler.Response
//	@Failure		500 {object} 	handler.Response
//	@Router		/v1/transactions [post]
func (p *Handler) CreateTransaction(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody CreateTransactionRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			handler.GenerateResponse(w, nil, err, http.StatusInternalServerError)
			return
		}

		_, err = p.AccountSvc.GetAccountByID(r.Context(), reqBody.AccountID)
		if err != nil {
			log.Err(err).Msg("error while creating the transaction")
			statusCode := http.StatusInternalServerError
			if errors.Is(err, account.ErrValidationFailed) {
				statusCode = http.StatusBadRequest
			} else if errors.Is(err, account.ErrAccountNotFound) {
				statusCode = http.StatusNotFound
			}
			handler.GenerateResponse(w, nil, err, statusCode)
			return
		}

		res, err := p.TxnSvc.CreateTransaction(r.Context(), &entity.Transaction{
			ID:              uuid.New().String(),
			AccountID:       reqBody.AccountID,
			OperationTypeID: reqBody.OperationTypeID,
			Amount:          reqBody.Amount,
		})
		if err != nil {
			statusCode := http.StatusInternalServerError
			log.Err(err).Msg("error while creating the transaction")
			if errors.Is(err, transaction.ErrValidationFailed) {
				statusCode = http.StatusBadRequest
			}
			handler.GenerateResponse(w, nil, err, statusCode)
			return
		}

		handler.GenerateResponse(w, CreateTransactionResponse{ID: res.ID}, nil, http.StatusCreated)
	}
}
