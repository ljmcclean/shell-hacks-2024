package main

import "net/http"

func NewServer(address string) *http.Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	return &http.Server{
		Addr:    address,
		Handler: mux,
	}
}
