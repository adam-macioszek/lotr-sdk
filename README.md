# Lord of the Rings SDK
This is a Go SDK for The One API, a Lord of the Rings API. It provides an lightweight Go wrapper around the /movie and /quote end points.
The One API can be found here: https://the-one-api.dev/.
## Installation

```
go get github.com/adam-macioszek/lotr-sdk
```
## Requirements
go 1.20
valid API key for The One API

## Usage
The SDK currently sources your API key from your local enviornment, on a mac or linux machine
this should do the trick:
```
export $LOTR_API_KEY=<your-api-key>
```

To use the SDK 
```
"github.com/adam-macioszek/lotr-sdk/pkg/movie"
"github.com/adam-macioszek/lotr-sdk/pkg/quote"

//Get All movies and print ID and Name for each movie
movies, err := movie.GetMovies()
	if err != nil {
		log.Println(err)
	}
	for _, movie := range movies {
		log.Println(movie.ID, movie.Name)
	}

//Get a quote by id and print the quote dialog
quoteId := "5cd96e05de30eff6ebccebb1"
testQuote, err := quote.GetQuoteByID(quoteId)
if err != nil {
    log.Println(err)
}
log.Println(testQuote.Dialog)
```

## Testing
- No real unit/end-to-end/integration testing was done, I merely validated the common use cases as you can see in cmd/main.go.

