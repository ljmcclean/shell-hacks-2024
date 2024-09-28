package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/server/auth"
	"github.com/ljmcclean/shell-hacks-2024/server/handlers"
	"github.com/ljmcclean/shell-hacks-2024/server/sessions"
	"github.com/ljmcclean/shell-hacks-2024/templates"
)

func addRoutes(mux *http.ServeMux, auth *auth.Authenticator) {
	assetsFS := http.FileServer(http.Dir("public"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	staticFS := http.FileServer(http.Dir("public/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFS))

	mux.Handle("/{$}", templ.Handler(templates.Index()))

	mux.Handle("/login", handlers.Login(auth))

	mux.Handle("/logout", sessions.IsAuthenticated(handlers.Logout()))

	mux.Handle("/callback", sessions.IsAuthenticated(handlers.Callback(auth)))

	mux.Handle("/dashboard", sessions.IsAuthenticated(handlers.Dashboard()))
}
