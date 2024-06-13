package service

import (
	"RyanFin/GoMountainWebScraper/src/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebScraper struct {
	Mountains []models.Mountain // many mountains to be added
}

// fetch random mountain data from online sources and store it in the mountain slice
// must filter out mountains that already exist in the MongoDB Atlas database
func (w *WebScraper) BuildMountain() error {
	// Get total list of mountains that already exist in the database

	currentMountains, err := fetchGoMountains()
	if err != nil {
		return err
	}

	// fmt.Println(currentMountains)

	for _, e := range currentMountains {
		fmt.Println(e.Name)
	}

	return nil
}

// Get total list of mountains that already exist in the database
func fetchGoMountains() ([]models.Mountain, error) {
	var mountains []models.Mountain

	// run API locally
	url := "http://localhost:8080/api/mountains"

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return mountains, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return mountains, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return mountains, err
	}

	err = json.Unmarshal([]byte(body), &mountains)
	if err != nil {
		fmt.Println(err)
		return mountains, err
	}

	return mountains, nil
}
