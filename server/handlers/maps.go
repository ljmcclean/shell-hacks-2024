package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/templates"
)

func SingleRoute() http.Handler {
	return templ.Handler(templates.SingleRoute())
}

func GroupRoute() http.Handler {
	return templ.Handler(templates.GroupRoute())
}

func getRouteData() {

}
