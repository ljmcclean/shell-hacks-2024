package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/ljmcclean/shell-hacks-2024/server/auth"
	"github.com/ljmcclean/shell-hacks-2024/server/sessions"
)

func Login(auth *auth.Authenticator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, err := sessions.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		session.Values["state"] = state

		if err := sessions.Store.Save(r, w, session); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	})
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
