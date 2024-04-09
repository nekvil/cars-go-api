package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/utils"
)

type ClientApiRepository struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClientApiRepository(baseURL string) *ClientApiRepository {
	return &ClientApiRepository{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (r *ClientApiRepository) GetByRegNum(regNum string) (model.Car, error) {
	encodedRegNum := url.QueryEscape(regNum)
	url := fmt.Sprintf("%s/info?regNum=%s", r.BaseURL, encodedRegNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Car{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return model.Car{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	utils.Logger.Debugf("Received response from external API with status code '%d'", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return model.Car{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var car model.Car
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return model.Car{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return car, nil
}
