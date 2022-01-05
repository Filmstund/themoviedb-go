package themoviedb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filmstund/themoviedb-go"
)

const configResponse = `{"images":{"base_url":"http://image.tmdb.org/t/p/","secure_base_url":"https://image.tmdb.org/t/p/","backdrop_sizes":["w300","w780","w1280","original"],"logo_sizes":["w45","w92","w154","w185","w300","w500","original"],"poster_sizes":["w92","w154","w185","w342","w500","w780","original"],"profile_sizes":["w45","w185","h632","original"],"still_sizes":["w92","w185","w300","original"]},"change_keys":["adult","air_date","also_known_as","alternative_titles","biography","birthday","budget","cast","certifications","character_names","created_by","crew","deathday","episode","episode_number","episode_run_time","freebase_id","freebase_mid","general","genres","guest_stars","homepage","images","imdb_id","languages","name","network","origin_country","original_name","original_title","overview","parts","place_of_birth","plot_keywords","production_code","production_companies","production_countries","releases","revenue","runtime","season","season_number","season_regular","spoken_languages","status","tagline","title","translations","tvdb_id","tvrage_id","type","video","videos"]}`

func TestClient_APIConfiguration(t *testing.T) {
	configCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configCalled = true
		if _, err := w.Write([]byte(configResponse)); err != nil {
			t.Fatalf("failed to write response: %b", err)
		}
	}))
	defer srv.Close()

	_, err := themoviedb.NewClient(
		context.Background(),
		"fake-key",
		themoviedb.WithClient(srv.Client()),
		themoviedb.WithBaseURL(srv.URL),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if configCalled == false {
		t.Errorf("configuration endpoint not called")
	}
}

func TestClient_APIConfiguration_provided(t *testing.T) {
	configCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configCalled = true
		if _, err := w.Write([]byte(configResponse)); err != nil {
			t.Fatalf("failed to write response: %b", err)
		}
	}))
	defer srv.Close()

	_, err := themoviedb.NewClient(
		context.Background(),
		"fake-key",
		themoviedb.WithClient(srv.Client()),
		themoviedb.WithBaseURL(srv.URL),
		themoviedb.WithAPIConfiguration(themoviedb.APIConfiguration{}),
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if configCalled == true {
		t.Errorf("configuration endpoint called, expected not to")
	}
}
