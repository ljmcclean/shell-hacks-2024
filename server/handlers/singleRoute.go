package handlers

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
	"os"
)

type RouteResponse struct {
	Routes []struct {
		Summary struct {
			Distance      float64 `json:"distance"`
			Duration      float64 `json:"duration"`
			transportType string  `json:"profile"`
			start         string  `json:"start"`
			end           string  `json:"end"`
		} `json:"summary"`
	} `json:"routes"`
}

func SingleRoute() http.Handler {
	data, err := getRouteData()
	if err != nil {
		return nil
	}
	duration := data.Routes[0].Summary.Duration
	distance := data.Routes[0].Summary.Distance
	transType := data.Routes[0].Summary.transportType
	start := data.Routes[0].Summary.start
	end := data.Routes[0].Summary.end

	return templ.Handler(singleRoute.SingleRoute(data))
}

func getRouteData() (*RouteResponse, error) {
	api_key := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://api.openrouteservice.org/v2/directions/%s?api_key=%s&start=%s&end=%s", transportType, api_key, start, end)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for a valid response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get valid response: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var routeResponse RouteResponse
	err = json.NewDecoder(resp.Body).Decode(&routeResponse)
	if err != nil {
		return nil, err
	}

	return &routeResponse, nil
}
