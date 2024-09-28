package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/templates"
	"github.com/ljmcclean/shell-hacks-2024/templates/components"
)

type route struct {
}

func SingleRoute() http.Handler {
	return templ.Handler(templates.SingleRoute())
}

func GroupRoute() http.Handler {
	return templ.Handler(templates.GroupRoute())
}

func GetMap(start, end string) {
	data := getRouteData(start, end)

	return templ.Handler(components.MapSection(data.duration, data.distance))
}

func getRouteData(start, end string) (*route, error) {
	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&key=%s", start, end, apiKey)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
