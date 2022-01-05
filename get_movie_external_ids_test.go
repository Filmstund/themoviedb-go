package themoviedb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filmstund/themoviedb-go"
)

const externalIDsResp = `{"id":634649,"imdb_id":"tt10872600","facebook_id":"SpiderManMovie","instagram_id":"spidermanmovie","twitter_id":"spidermanmovie"}`

func TestClient_MovieExternalIDs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(externalIDsResp)); err != nil {
			t.Fatalf("failed to write response: %b", err)
		}
	}))
	defer srv.Close()

	client, err := themoviedb.NewClient(
		context.Background(),
		"fake-api-key",
		themoviedb.WithAPIConfiguration(themoviedb.APIConfiguration{}),
		themoviedb.WithBaseURL(srv.URL),
		themoviedb.WithClient(srv.Client()),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	exts, err := client.MovieExternalIDs(context.Background(), 634649)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if exts.ID != 634649 {
		t.Errorf("wrong movie ID")
	}
	if *exts.ImdbID != "tt10872600" {
		t.Errorf("wrong movie imdbID")
	}
	if *exts.FacebookID != "SpiderManMovie" {
		t.Errorf("wrong movie facebook ID")
	}
	if *exts.InstagramID != "spidermanmovie" || *exts.TwitterID != "spidermanmovie" {
		t.Errorf("wrong movie twitter/instagram ID")
	}
}
