package handlers

import (
	"net/http"
	"github.com/auth0-community/auth0"
	"gopkg.in/square/go-jose.v2"
	"log"
	"github.com/nestorsokil/goproc-img/api-ms/config"
)

func doAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(config.Settings.PrivateKey)
		secretProvider := auth0.NewKeyProvider(secret)
		audience := config.Settings.Audience
		configuration := auth0.NewConfiguration(
			secretProvider,
			audience,
			config.Settings.Auth0DomainName,
			jose.HS256)
		validator := auth0.NewValidator(configuration)

		token, err := validator.ValidateRequest(r)
		if err != nil {
			log.Println("[WARN] Token is not valid:", token, err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
