package themoviedb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type MovieSearchResult struct {
	Page    int `json:"page"`
	Results []struct {
		ID               int     `json:"id"`
		Adult            bool    `json:"adult"`
		BackdropPath     *string `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func (client *Client) SearchMovie(ctx context.Context, query string, year int) (*MovieSearchResult, error) {
	params := url.Values{}
	params.Set("api_key", client.apiKey)
	params.Set("query", query)
	params.Set("year", strconv.Itoa(year))
	searchURL := fmt.Sprintf("%s/search/movie?%s", client.baseURL, params.Encode())

	results := new(MovieSearchResult)
	if err := client.get(ctx, searchURL, results); err != nil {
		return nil, fmt.Errorf("searchMovie: %w", err)
	}
	return results, nil
}
