package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/auth/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/auth/pkg/service"
	"log"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

func (h AuthHandler) NotImplementedHandler(w http.ResponseWriter, _ *http.Request) {
	writeResponse(w, http.StatusOK, "Handler not implemented...")
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.Service.Login(loginRequest)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, err.Error())
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

/*
	Verify Sample URL string
 	http://localhost:5001/auth/verify?token=somevalidtokenstring&routeName=GetCustomer&customer_id=1&account_id=1
*/
func (h AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	// converting from Query to map type
	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		isAuthorized, appError := h.Service.Verify(urlParams)
		if appError != nil {
			writeResponse(w, http.StatusForbidden, notAuthorizedResponse())
		} else {
			if isAuthorized {
				writeResponse(w, http.StatusOK, authorizedResponse())
			} else {
				writeResponse(w, http.StatusForbidden, notAuthorizedResponse())
			}
		}
	} else {
		writeResponse(w, http.StatusForbidden, "missing token")
	}
}

func notAuthorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": false}
}

func authorizedResponse() map[string]bool {
	return map[string]bool{"isAuthorized": true}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
