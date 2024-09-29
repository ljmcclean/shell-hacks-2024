package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/ljmcclean/shell-hacks-2024/templates"
	"github.com/ljmcclean/shell-hacks-2024/templates/components"
)

type ORSResponse struct {
	Routes []struct {
		Summary struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
		} `json:"summary"`
	} `json:"routes"`
}

func SingleRoute() http.Handler {
	return templ.Handler(templates.SingleRoute())
}

func GroupRoute() http.Handler {
	return templ.Handler(templates.GroupRoute())
}

func Map() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %v", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		start, end := r.FormValue("start"), r.FormValue("end")
		start, end = strings.ReplaceAll(start, " ", ""), strings.ReplaceAll(end, " ", "")

		distance, duration := getRouteData(start, end)

		templ.Handler(components.MapSection(distance, duration)).ServeHTTP(w, r)
	})
}

func getRouteData(start, end string) (float64, float64) {
	apiKey := os.Getenv("ORS_API_KEY")
	url := fmt.Sprintf("https://api.openrouteservice.org/v2/directions/driving-car?api_key=%s&start=%s&end=%s", apiKey, start, end)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var orsResp ORSResponse
	if err := json.Unmarshal(body, &orsResp); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	if len(orsResp.Routes) > 0 {
		summary := orsResp.Routes[0].Summary
		return summary.Distance, summary.Duration
	} else {
		log.Println("No routes found.")
		return 0.0, 0.0
	}
}
