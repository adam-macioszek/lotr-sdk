package quote

import (
	"testing"

	"github.com/adam-macioszek/lotr-sdk/quote"
)

func TestGetQuoteById(t *testing.T) {
	quoteId := "5cd96e05de30eff6ebccebb1"
	quoteDialog := "Hey, stinker, don't go gettingtoo far ahead."
	got, err := quote.GetQuoteByID(quoteId)
	want := quoteDialog

	if err != nil {
		t.Errorf("encountered error %err", err)
	}
	if got.Dialog != want {
		t.Errorf("got %q, wanted %q", got.Dialog, want)
	}
}
