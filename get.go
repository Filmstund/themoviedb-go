package themoviedb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) get(ctx context.Context, url string, entity interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return newAPIError(resp.Body, "get: non-ok response: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(entity); err != nil {
		return fmt.Errorf("get: %w", err)
	}
	return nil
}
