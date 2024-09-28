package handlers

import (
	"net/http"
)

func Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dashboard"))
	})
}
