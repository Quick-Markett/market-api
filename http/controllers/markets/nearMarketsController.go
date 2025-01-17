package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main.go/models"
)

type LocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func FindMarketsWithinRadius(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	var locReq LocationRequest
	if err := c.ShouldBindJSON(&locReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	apiKey := os.Getenv("GOOGLE_PLACES_API_KEY")
	if apiKey == "" {
		log.Printf("API key is missing")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is missing"})
		return
	}

	location := fmt.Sprintf("%f,%f", locReq.Latitude, locReq.Longitude)
	radius := "5000" // 5km
	placeType := "market"
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&radius=%s&type=%s&key=%s", location, radius, placeType, apiKey)

	log.Printf("key: %s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from Google Places API"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	var placesResponse struct {
		Results []struct {
			Name     string `json:"name"`
			Vicinity string `json:"vicinity"`
			Geometry struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
		} `json:"results"`
	}

	if err := json.Unmarshal(body, &placesResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	NearMarkets := []models.NearMarkets{}

	for _, result := range placesResponse.Results {
		NearMarkets = append(NearMarkets, models.NearMarkets{
			Name:    result.Name,
			Address: result.Vicinity,
			Location: struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			}{
				Latitude:  result.Geometry.Location.Lat,
				Longitude: result.Geometry.Location.Lng,
			},
		})
	}

	c.JSON(http.StatusOK, NearMarkets)
}
