package themoviedb

import (
	"context"
	"fmt"
)

type MovieExternalIDs struct {
	ID          int64   `json:"id"`
	ImdbID      *string `json:"imdb_id"`
	FacebookID  *string `json:"facebook_id"`
	InstagramID *string `json:"instagram_id"`
	TwitterID   *string `json:"twitter_id"`
}

func (client *Client) MovieExternalIDs(ctx context.Context, movieID int64) (*MovieExternalIDs, error) {
	url := fmt.Sprintf("%s/movie/%d/external_ids?api_key=%s", client.baseURL, movieID, client.apiKey)
	external := new(MovieExternalIDs)
	if err := client.get(ctx, url, external); err != nil {
		return nil, fmt.Errorf("movieExternalIDs: %w", err)
	}
	return external, nil
}
