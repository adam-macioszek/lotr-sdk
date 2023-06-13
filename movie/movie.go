package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/adam-macioszek/lotr-sdk/request"
)

type movieResponse struct {
	Movies []Movie `json:"docs"`
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Page   int     `json:"page"`
	Pages  int     `json:"pages"`
}

type Movie struct {
	ID                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           float32 `json:"runtimeInMinutes"`
	BudgetInMillions           float32 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float32 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float32 `json:"rottenTomatoesScore"`
}

func GetMovies() ([]Movie, error) {
	response, err := request.CallApi("/movie")
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return []Movie{}, err
	}
	movies, err := deserializeMovies(response)
	if err != nil {
		return []Movie{}, err
	}
	return movies, nil
}
func GetMovieByID(id string) (Movie, error) {
	response, err := request.CallApi("/movie/" + id)
	if err != nil {
		log.Printf("error calling API %v\n", err)
		return Movie{}, err
	}
	movies, err := deserializeMovies(response)
	if err != nil {
		return Movie{}, err
	}
	if len(movies) > 0 {
		return movies[0], nil
	}

	return Movie{}, errors.New("malformed API response")
}
func GetMovieByName(movieName string) (Movie, error) {
	movies, err := GetMovies()
	if err != nil {
		return Movie{}, err
	}
	for _, movie := range movies {
		if movie.Name == movieName {
			return movie, nil
		}
	}
	errorMessage := fmt.Sprintf("Cannont find matching name for %s in Lord of the Rings movies", movieName)
	return Movie{}, errors.New(errorMessage)
}
func GetMovieNames() ([]string, error) {
	var result []string
	movies, err := GetMovies()
	if err != nil {
		return []string{}, err
	}
	for _, movie := range movies {
		result = append(result, movie.Name)
	}
	return result, nil
}

func deserializeMovies(responceBytes []byte) ([]Movie, error) {
	var movieResp movieResponse
	if err := json.Unmarshal(responceBytes, &movieResp); err != nil {
		log.Printf("error deserializing json data %v\n", err)
		return []Movie{}, err
	}
	return movieResp.Movies, nil
}
