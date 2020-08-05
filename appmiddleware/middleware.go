package appmiddleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("CHECKID")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Set CHECKID cookies", http.StatusUnauthorized)
			return
		}

		correctValue := "let-me-pass"
		if correctValue != cookie.Value {
			http.Error(w, "Incorrect CHECKID cookies", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
