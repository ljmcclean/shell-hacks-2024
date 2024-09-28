package server

import (
	"net/http"

	"github.com/ljmcclean/shell-hacks-2024/server/auth"
)

func New(address string, auth *auth.Authenticator) *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux, auth)

	return &http.Server{
		Addr:    address,
		Handler: mux,
	}
}
