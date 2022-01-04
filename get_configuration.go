package themoviedb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type APIConfiguration struct {
	Images struct {
		BaseUrl       string   `json:"base_url"`
		SecureBaseUrl string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

// APIConfiguration requests the current API config from the API.
func (client *Client) APIConfiguration(ctx context.Context) (*APIConfiguration, error) {
	url := fmt.Sprintf("%s/configuration?api_key=%s", client.baseURL, client.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("apiConfiguration: %w", err)
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("apiConfiguration: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, newAPIError(resp.Body, "apiConfiguration: non-ok response: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	apiConfig := new(APIConfiguration)
	if err := decoder.Decode(apiConfig); err != nil {
		return nil, fmt.Errorf("apiConfiguration: %w", err)
	}
	return apiConfig, nil
}
