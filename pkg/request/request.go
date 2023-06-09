package request

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const API_ENDPOINT string = "https://the-one-api.dev/v2"

func CallApi(uri string) ([]byte, error) {
	req, err := buildRequest(uri)
	if err != nil {
		log.Printf("error building HTTP request: %v\n", err)
		return []byte{}, err
	}

	responseBytes, err := makeRequest(req)
	if err != nil {
		log.Printf("error with HTTP response: %v\n", err)
		return []byte{}, err
	}

	return responseBytes, nil

}

func buildRequest(uri string) (*http.Request, error) {

	request, err := http.NewRequest(
		http.MethodGet,
		API_ENDPOINT+uri,
		nil,
	)
	if err != nil {
		log.Printf("error creating HTTP request: %v", err)
		return request, err
	}

	if os.Getenv("LOTR_API_KEY") == "" {
		return request, errors.New("API token not set. Export it locally.")
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("LOTR_API_KEY"))
	return request, nil
}

func makeRequest(request *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("error sending HTTP request: %v", err)
		return []byte{}, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		return []byte{}, errors.New("unauthorized request, check your API token")
	}
	if res.Header.Get("X-Ratelimit-Remaining") == "0" {
		return []byte{}, errors.New("api rate limit reached, rate limit will reset within 10 minutes")
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading HTTP response body: %v\n", err)
		return responseBytes, err
	}

	return responseBytes, nil
}
