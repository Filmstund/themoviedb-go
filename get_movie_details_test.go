package themoviedb_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filmstund/themoviedb-go"
)

const movieDetailsResp = `{"adult":false,"backdrop_path":"/lvOLivVeX3DVVcwfVkxKf0R22D8.jpg","belongs_to_collection":{"id":1241,"name":"Harry Potter Collection","poster_path":"/eVPs2Y0LyvTLZn6AP5Z6O2rtiGB.jpg","backdrop_path":"/wfnMt6LGqYHcNyOfsuusw5lX3bL.jpg"},"budget":125000000,"genres":[{"id":12,"name":"Adventure"},{"id":14,"name":"Fantasy"}],"homepage":"https://www.warnerbros.com/movies/harry-potter-and-sorcerers-stone/","id":671,"imdb_id":"tt0241527","original_language":"en","original_title":"Harry Potter and the Philosopher's Stone","overview":"Harry Potter has lived under the stairs at his aunt and uncle's house his whole life. But on his 11th birthday, he learns he's a powerful wizard—with a place waiting for him at the Hogwarts School of Witchcraft and Wizardry. As he learns to harness his newfound powers with the help of the school's kindly headmaster, Harry uncovers the truth about his parents' deaths—and about the villain who's to blame.","popularity":428.246,"poster_path":"/wuMc08IPKEatf9rnMNXvIDxqP4W.jpg","production_companies":[{"id":174,"logo_path":"/IuAlhI9eVC9Z8UQWOIDdWRKSEJ.png","name":"Warner Bros. Pictures","origin_country":"US"},{"id":437,"logo_path":"/nu20mtwbEIhUNnQ5NXVhHsNknZj.png","name":"Heyday Films","origin_country":"GB"},{"id":436,"logo_path":"/A7WCkG3F0NFvjGCwUnclpGdIu9E.png","name":"1492 Pictures","origin_country":"US"}],"production_countries":[{"iso_3166_1":"GB","name":"United Kingdom"},{"iso_3166_1":"US","name":"United States of America"}],"release_date":"2001-11-16","revenue":976475550,"runtime":152,"spoken_languages":[{"english_name":"English","iso_639_1":"en","name":"English"}],"status":"Released","tagline":"Let the Magic Begin.","title":"Harry Potter and the Philosopher's Stone","video":false,"vote_average":7.9,"vote_count":21778}`

func TestClient_MovieDetails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(movieDetailsResp)); err != nil {
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

	deets, err := client.MovieDetails(context.Background(), 671)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if deets.ID != 671 {
		t.Errorf("wrong movie ID")
	}
	if deets.Title != "Harry Potter and the Philosopher's Stone" {
		t.Errorf("wrong movie title")
	}
	if deets.Popularity != 428.246 {
		t.Errorf("wrong movie popularity")
	}
}
