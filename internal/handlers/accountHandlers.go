package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AccountHandlers struct {
	Service service.AccountService
	L       *log.Logger
}

func (ah AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		renderJSON(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := ah.Service.NewAccount(request)
		if appError != nil {
			renderJSON(w, appError.StatusCode, appError.AsMessage())
		} else {
			renderJSON(w, http.StatusCreated, account)
		}
	}
}
