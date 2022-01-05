package themoviedb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filmstund/themoviedb-go"
)

const movieResultsResp = `{"page":1,"results":[{"adult":false,"backdrop_path":"/1Rr5SrvHxMXHu5RjKpaMba8VTzi.jpg","genre_ids":[28,12,878],"id":634649,"original_language":"en","original_title":"Spider-Man: No Way Home","overview":"Peter Parker is unmasked and no longer able to separate his normal life from the high-stakes of being a super-hero. When he asks for help from Doctor Strange the stakes become even more dangerous, forcing him to discover what it truly means to be Spider-Man.","popularity":8817.063,"poster_path":"/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg","release_date":"2021-12-15","title":"Spider-Man: No Way Home","video":false,"vote_average":8.4,"vote_count":3427},{"adult":false,"backdrop_path":"/9nuSlccCMWtIejzIn6kpS4zlWk5.jpg","genre_ids":[99,10770],"id":923349,"original_language":"fr","original_title":"StreamVF - Spider-Man No Way Home","overview":"Explore the multiverse with the French voice actors of all three Spider-Men!","popularity":0,"poster_path":"/610ij9wU2V1DjuZ5KPn92hBrcKP.jpg","release_date":"2021-12-23","title":"StreamVF - Spider-Man No Way Home","video":false,"vote_average":0,"vote_count":0}],"total_pages":1,"total_results":2}`

func TestClient_SearchMovie(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(movieResultsResp)); err != nil {
			t.Fatalf("failed to write response: %b", err)
		}
	}))
	defer srv.Close()

	client, err := themoviedb.NewClient(
		context.Background(),
		"47dd57b1d0383ef68392322a04f263a0",
		themoviedb.WithAPIConfiguration(themoviedb.APIConfiguration{}),
		themoviedb.WithClient(srv.Client()),
		themoviedb.WithBaseURL(srv.URL),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	results, err := client.SearchMovie(context.Background(), "Spider-Man: No Way Home", 2021)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if results.TotalResults != 2 || len(results.Results) != 2 {
		t.Errorf("expected 2 results, got %d (%d)", results.TotalResults, len(results.Results))
	}
	if results.Page != 1 {
		t.Errorf("expected 1 page, got %d", results.Page)
	}
	if results.TotalPages != 1 {
		t.Errorf("expected 1 total page, got %d", results.TotalPages)
	}
	if results.Results[0].Title != "Spider-Man: No Way Home" {
		t.Errorf("expected 'Spider-Man: No Way Home', got %s", results.Results[0].Title)
	}
}
