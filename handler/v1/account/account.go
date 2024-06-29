package account

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hax0rr/transaction-service/handler"
	"github.com/hax0rr/transaction-service/internal/entity"
	"github.com/hax0rr/transaction-service/internal/service/account"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"net/http"
)

type Handler struct {
	service account.IService
}

func NewHandler(service account.IService) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAccount
//
//	@Summary	Create a new account
//	@Tags		Accounts
//	@Accept		json
//	@Produce	json
//	@Param		account			body		CreateAccountRequest	true	"Create Account"
//	@Success	201				{object}	handler.Response
//	@Failure		500 			{object} 	handler.Response
//	@Failure		400 			{object} 	handler.Response
//	@Router		/v1/accounts 	[post]
func (p *Handler) CreateAccount(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody CreateAccountRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			handler.GenerateResponse(w, nil, err, http.StatusInternalServerError)
			return
		}

		res, err := p.service.CreateAccount(r.Context(), &entity.Account{
			ID:             uuid.New().String(),
			DocumentNumber: reqBody.DocumentNumber,
		})
		if err != nil {
			log.Err(err).Msg("error while creating the account")
			errCode := http.StatusInternalServerError
			if errors.Is(err, account.ErrValidationFailed) {
				errCode = http.StatusBadRequest
			}
			handler.GenerateResponse(w, nil, err, errCode)
			return
		}

		handler.GenerateResponse(w, CreateAccountResponse{ID: res.ID}, nil, http.StatusCreated)
	}
}

// GetAccountByID
//
//	@Summary	Get account by id
//	@Tags		Accounts
//	@Accept		json
//	@Produce	json
//	@Param		account_id		path		string	true	"account id"
//	@Success	200				{object}	handler.Response
//	@Failure		404 {object} 	handler.Response
//	@Failure		400 {object} 	handler.Response
//	@Failure		500 {object} 	handler.Response
//	@Router		/v1/accounts/{account_id} [get]
func (p *Handler) GetAccountByID(log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accountID := vars["accountID"]

		res, err := p.service.GetAccountByID(r.Context(), accountID)
		if err != nil {
			log.Err(err).Msg("error while getting the account details")
			errCode := http.StatusInternalServerError
			if errors.Is(err, account.ErrValidationFailed) {
				errCode = http.StatusBadRequest
			} else if errors.Is(err, account.ErrAccountNotFound) {
				errCode = http.StatusNotFound
			}
			handler.GenerateResponse(w, nil, err, errCode)
			return
		}

		handler.GenerateResponse(w, GetAccountByIDResponse{ID: res.ID, DocumentNumber: res.DocumentNumber}, nil, http.StatusOK)
	}
}
