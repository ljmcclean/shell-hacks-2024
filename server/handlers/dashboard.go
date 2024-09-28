package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/templates"
)

func Dashboard() http.Handler {
	return templ.Handler(templates.Dashboard())
}
