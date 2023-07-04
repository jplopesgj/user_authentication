package middleware

import "net/http"

func ContenTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "applicaton-json")
		next.ServeHTTP(w, r)
	})

}
