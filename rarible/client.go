package rarible

import (
	"bytes"
	"encoding/json"
	"fmt"
	"l0s0s/Rarible-API/model"
	"net/http"
)

const (
	rariblePath string = "https://api.rarible.org/v0.1"
)

func NewClient(apiKey, referer string) Client {
	return Client{
		apiKey:  apiKey,
		referer: referer,
	}
}

type Client struct {
	apiKey  string
	referer string
}

func (cli *Client) GetNFTOwnership(id string) (model.Ownership, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ownerships/%s", rariblePath, id), nil)
	if err != nil {
		return model.Ownership{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("X-API-KEY", cli.apiKey)
	req.Header.Add("Referer", cli.referer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Ownership{}, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		response := struct {
			Message string `json:"message"`
		}{}

		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return model.Ownership{}, fmt.Errorf("failed to decode response: %w", err)
		}

		return model.Ownership{}, fmt.Errorf("%s", response.Message)
	}

	response := model.Ownership{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.Ownership{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}

func (cli *Client) GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
	requestBody := map[string]interface{}{
		"properties":   properties,
		"collectionId": collectionID,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return model.TraitsRarity{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items/traits/rarity", rariblePath), bytes.NewBuffer(jsonBody))
	if err != nil {
		return model.TraitsRarity{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("X-API-KEY", cli.apiKey)
	req.Header.Add("Referer", cli.referer)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.TraitsRarity{}, fmt.Errorf("failed to execute request: %w", err)
	}

	fmt.Println("Response Status Code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		response := struct {
			Message string `json:"message"`
		}{}

		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return model.TraitsRarity{}, fmt.Errorf("failed to decode response: %w", err)
		}

		return model.TraitsRarity{}, fmt.Errorf("%s", response.Message)
	}

	response := model.TraitsRarity{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.TraitsRarity{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
