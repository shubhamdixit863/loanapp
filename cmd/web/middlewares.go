package main

import (
	"customerservice/internal/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		} else {

			_, err := utils.ValidateToken(header)
			if err != nil {
				json.NewEncoder(w).Encode("Not a Valid token")
				return
			}

			next.ServeHTTP(w, r)
		}

	})
}
