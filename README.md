<p align="center">
  <img src=".github/assets/gopher-movie.svg" alt="A gopher eating popcorn while watching a movie" width="160">
</p>

# The Movie DB Go library

[![Go Reference](https://pkg.go.dev/badge/github.com/filmstund/themoviedb-go.svg)](https://pkg.go.dev/github.com/filmstund/themoviedb-go)
[![Sanity checks](https://github.com/filmstund/themoviedb-go/actions/workflows/checks.yml/badge.svg)](https://github.com/filmstund/themoviedb-go/actions/workflows/checks.yml)

A small, dependency-free Go client for [The Movie Database (TMDB) API v3](https://developer.themoviedb.org/docs).

It covers a handful of endpoints rather than the whole API:

| Method | TMDB endpoint |
| --- | --- |
| `GetAPIConfiguration` | `GET /configuration` |
| `SearchMovie` | `GET /search/movie` |
| `MovieDetails` | `GET /movie/{movie_id}` |
| `MovieExternalIDs` | `GET /movie/{movie_id}/external_ids` |

## Installation

```bash
go get github.com/filmstund/themoviedb-go
```

Requires Go 1.27 or newer.

## Usage

You need a TMDB **API key (v3 auth)**, which you can get from your [TMDB account settings](https://www.themoviedb.org/settings/api).

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/filmstund/themoviedb-go"
)

func main() {
	ctx := context.Background()

	client, err := themoviedb.NewClient(ctx, os.Getenv("TMDB_API_KEY"),
		themoviedb.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	results, err := client.SearchMovie(ctx, "Spider-Man: No Way Home", 2021)
	if err != nil {
		log.Fatal(err)
	}
	if len(results.Results) == 0 {
		log.Fatal("no matches")
	}

	movieID := results.Results[0].ID

	details, err := client.MovieDetails(ctx, movieID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (%s), %d min\n", details.Title, details.ReleaseDate, details.Runtime)

	ids, err := client.MovieExternalIDs(ctx, movieID)
	if err != nil {
		log.Fatal(err)
	}
	if ids.ImdbID != nil {
		fmt.Println("IMDb:", *ids.ImdbID)
	}
}
```

### API configuration and image URLs

`NewClient` calls `/configuration` once on creation and caches the result. Read it with `client.APIConfiguration()`, for example to build image URLs:

```go
cfg := client.APIConfiguration()
posterURL := cfg.Images.SecureBaseURL + "w500" + details.PosterPath
```

To skip that request (in tests, or if you already have the configuration), pass `WithAPIConfiguration`.

### Options

| Option | Description |
| --- | --- |
| `WithTimeout(d time.Duration)` | Timeout for every HTTP request. Zero means no timeout. |
| `WithClient(c *http.Client)` | Use your own `*http.Client` (transport, proxies, retries, etc.). |
| `WithBaseURL(url string)` | Override the API base URL (defaults to `https://api.themoviedb.org/3`). |
| `WithAPIConfiguration(cfg APIConfiguration)` | Use the given configuration instead of fetching it in `NewClient`. |

Note that `WithClient` replaces the HTTP client, so apply `WithTimeout` after it if you use both.

### Errors

When TMDB responds with a non-200 status, the returned error wraps an `*APIError`. Its `Response()` holds the decoded TMDB error body:

```go
var apiErr *themoviedb.APIError
if errors.As(err, &apiErr) {
	resp := apiErr.Response()
	log.Printf("TMDB error %d: %s", resp.StatusCode, resp.StatusMessage)
}
```

## Testing against a fake server

Point the client at an `httptest.Server` with `WithBaseURL` and `WithClient`, and pass `WithAPIConfiguration` so `NewClient` doesn't call out:

```go
srv := httptest.NewServer(handler)
defer srv.Close()

client, err := themoviedb.NewClient(t.Context(), "test-key",
	themoviedb.WithAPIConfiguration(themoviedb.APIConfiguration{}),
	themoviedb.WithClient(srv.Client()),
	themoviedb.WithBaseURL(srv.URL),
)
```

## Development

The `Makefile` runs the same checks as CI. You need [golangci-lint](https://golangci-lint.run/) v2 installed.

```bash
make        # tidy, format, lint, test, build, and check for an unchanged tree
make test   # tests only
make lint   # golangci-lint only
```

Releases are cut automatically from `master` by [go-semantic-release](https://github.com/go-semantic-release/semantic-release), so write commit messages as [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `chore:`, …).

## Attribution

This product uses the TMDB API but is not endorsed or certified by TMDB.

## License

[MIT](LICENSE)
