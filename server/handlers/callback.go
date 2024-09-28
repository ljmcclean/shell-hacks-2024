package handlers

import (
	"net/http"

	"github.com/ljmcclean/shell-hacks-2024/server/auth"
	"github.com/ljmcclean/shell-hacks-2024/server/sessions"
)

func Callback(auth *auth.Authenticator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessions.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, "Failed to retrieve session", http.StatusUnauthorized)
			return
		}

		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
			return
		}

		token, err := auth.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange an authorization code for a token", http.StatusUnauthorized)
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), *token)
		if err != nil {
			http.Error(w, "Failed to verify ID Token", http.StatusInternalServerError)
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["access_token"] = token.AccessToken
		session.Values["profile"] = profile

		if err := sessions.Store.Save(r, w, session); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
	})
}
