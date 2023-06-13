package movie

import (
	"testing"

	"github.com/adam-macioszek/lotr-sdk/movie"
)

func TestGetMovieById(t *testing.T) {
	movieID := "5cd95395de30eff6ebccde5b"
	movieName := "The Two Towers"
	got, err := movie.GetMovieByID(movieID)
	want := movieName

	if err != nil {
		t.Errorf("encountered error %err", err)
	}
	if got.Name != want {
		t.Errorf("got %q, wanted %q", got.Name, want)
	}
}
