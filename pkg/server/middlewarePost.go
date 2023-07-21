package server

import (
	"encoding/json"
	"flights/pkg/types"
	"fmt"
	"log"
	"net/http"
)

// MiddlewarePost Middleware function for post method only
func MiddlewarePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if r.Method != http.MethodPost {

			var rest types.RestFul
			rest.AddError(fmt.Errorf("Method not allowed"))
			w.WriteHeader(http.StatusMethodNotAllowed)
			err := json.NewEncoder(w).Encode(rest)
			if err != nil {
				log.Printf("myMiddlewarePost().json.NewEncoder(w).Encode(rest).Error: %v", err)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
