package sessions

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("secret"))

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "auth-session")
		r = r.WithContext(context.WithValue(r.Context(), "session", session))
		next.ServeHTTP(w, r)
	})
}
