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

func getRouteData(start, end string) {
	apiKey, err := 

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&key=%s", start, end, apiKey)
	res, err := http.Get()

}
