package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/server/sessions"
	"github.com/ljmcclean/shell-hacks-2024/templates"
)

func Dashboard() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessions.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, "Failed to retrieve session", http.StatusUnauthorized)
			return
		}

		profile, ok := session.Values["profile"].(map[string]interface{})
		if !ok {
			http.Error(w, "Profile not found", http.StatusNotFound)
			return
		}

		userName, ok := profile["name"].(string)
		if !ok {
			userName = "Guest"
		}

		templ.Handler(templates.Dashboard(userName)).ServeHTTP(w, r)
	})
}
