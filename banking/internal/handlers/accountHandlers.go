package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/lib"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AccountHandlers struct {
	Service service.AccountService
	L       *log.Logger
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.RenderJSON(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.Service.NewAccount(request)
		if appError != nil {
			lib.RenderJSON(w, appError.StatusCode, appError.AsMessage())
		} else {
			lib.RenderJSON(w, http.StatusCreated, account)
		}
	}
}

// MakeTransaction endpoint customers/2001/account/95470
func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		lib.RenderJSON(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.Service.MakeTransaction(request)

		if appError != nil {
			lib.RenderJSON(w, appError.StatusCode, appError.AsMessage())
		} else {
			lib.RenderJSON(w, http.StatusOK, account)
		}
	}
}
