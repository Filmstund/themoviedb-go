package themoviedb

import (
	"context"
	"fmt"
)

type APIConfiguration struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

// GetAPIConfiguration requests the current API config from the API.
func (client *Client) GetAPIConfiguration(ctx context.Context) (*APIConfiguration, error) {
	url := fmt.Sprintf("%s/configuration?api_key=%s", client.baseURL, client.apiKey)
	apiConfig := new(APIConfiguration)
	if err := client.get(ctx, url, apiConfig); err != nil {
		return nil, fmt.Errorf("apiConfiguration: %w", err)
	}
	return apiConfig, nil
}
