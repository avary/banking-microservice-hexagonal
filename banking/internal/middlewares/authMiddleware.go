package middlewares

import (
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/errs"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/lib"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Auth struct {
	Repo domain.AuthRepository
}

func (a Auth) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.Repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{Message: "Unauthorized", StatusCode: http.StatusForbidden}
					lib.RenderJSON(w, appError.StatusCode, appError.AsMessage())
				}
			} else {
				lib.RenderJSON(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
