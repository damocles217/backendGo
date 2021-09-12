package middlewares

import (
	"net/http"

	"github.com/damocles217/server/router/user/config"
	"github.com/gorilla/mux"
)

func AuthMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("c_user")

			if err != nil {
				h.ServeHTTP(w, r)
				return
			}

			value, err := config.AesDecrypt([]byte(cookie.Value), []byte("0123456789abcdef"))

			if err != nil {
				h.ServeHTTP(w, r)
				return
			}

			if string(value) == "" {
				h.ServeHTTP(w, r)
				return
			}

			r.Header.Set("Authorized", "true")

			h.ServeHTTP(w, r)
		})
	}
}
