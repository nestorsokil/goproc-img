package handlers

import (
	"net/http"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/nestorsokil/goproc-img/api-ms/util"
	"strings"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/nestorsokil/goproc-img/api-ms/config"
	"io"
)


type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func doLogin(response http.ResponseWriter, request *http.Request) {
	var user UserCredentials
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		util.Respond403(response, "No credentials provided")
		return
	}

	//FIXME: stub auth
	if strings.ToLower(user.Username) != "nsokil" {
		if user.Password != "test123" {
			util.Respond403(response, "Invalid credentials")
			return
		}
	}

	claims := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(8*time.Hour).Unix(),
		Issuer: "admin",
	}
	signer := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := signer.SignedString(config.Settings.PrivateKeyPath)
	if err != nil {
		util.Respond500(response, "Error creating a token string")
		log.Println("[ERROR] Error creating a token string", err)
	}

	util.RespondJson(response, TokenResponse{tokenString})
}

func doLoginStub(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Authorization disabled.")
}

func doAuth(next http.HandlerFunc) (http.HandlerFunc) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return config.Settings.PublicKeyPath, nil
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Get("Authorization")
		jwtToken, err := jwtrequest.ParseFromRequest(r, jwtrequest.OAuth2Extractor, keyFunc)
		if err != nil || !jwtToken.Valid {
			util.Respond403(w, "Wrong credentials.")
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handler)
}
