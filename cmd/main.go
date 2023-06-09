package main

import (
	"log"

	"github.com/adam-macioszek/lotr-sdk/pkg/movie"
	"github.com/adam-macioszek/lotr-sdk/pkg/quote"
)

const API_ENDPOINT string = "https://the-one-api.dev/v2"

func main() {
}

func testMovies() {
	movies, err := movie.GetMovies()
	if err != nil {
		log.Println(err)
	}
	for _, movie := range movies {
		log.Println(movie.ID, movie.Name)
	}

	TwoTowers, err := movie.GetMovieByID("5cd95395de30eff6ebccde5b")
	if err != nil {
		log.Println(err)
	}
	log.Println(TwoTowers.Name, TwoTowers.BudgetInMillions)
	ReturnOftheKing, err := movie.GetMovieByName("The Return of the King")
	if err != nil {
		log.Println(err)
	}
	log.Println(ReturnOftheKing.Name, ReturnOftheKing.AcademyAwardWins)
}

func testQuotes() {
	//Getting all quotes
	quotes, err := quote.GetQuotes()
	if err != nil {
		log.Println(err)
	}
	for _, quote := range quotes {
		log.Println(quote.ID, quote.Dialog)
	}

	//Getting quotes from one movie, i.e TwinTowers
	movieId := "5cd95395de30eff6ebccde5b"
	testQuotes, err := quote.GetQuotesFromMovie(movieId)
	if err != nil {
		log.Println(err)
	}
	for _, quote := range testQuotes {
		log.Println(quote.ID, quote.Dialog)
	}

	//Get Quote by ID
	quoteId := "5cd96e05de30eff6ebccebb1"
	testQuote, err := quote.GetQuoteByID(quoteId)
	if err != nil {
		log.Println(err)
	}
	log.Println(testQuote.Dialog, testQuote.ID)
	//quoteDialog := "Hey, stinker, don't go gettingtoo far ahead."

}
