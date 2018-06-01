package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeppyS/what-to-watch/proto"
)

const (
	movieIndex   = "movies"
	movieMapping = "movie"
)

type ElasticsearchClient struct {
	httpClient *http.Client
	url        string
}

func NewElasticsearchClient(c *http.Client, url string) *ElasticsearchClient {
	return &ElasticsearchClient{c, url}
}

func (s *MovieService) AutocompleteMovies(text string) ([]*proto.MoviesList_Movie, error) {
	return nil, nil
}

func (c *ElasticsearchClient) ClearMovieIndex() error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s/%s", c.url, movieIndex), nil)
	if err != nil {
		return fmt.Errorf("Failed creating request: %v", err)
	}

	_, err = c.httpClient.Do(req)

	return err
}

func (c *ElasticsearchClient) BulkIndexMovies(movies []*proto.MoviesList_Movie) error {
	err := c.ClearMovieIndex()
	if err != nil {
		return fmt.Errorf("Problem clearing movie index: %v", err)
	}

	for _, movie := range movies {
		b, err := json.Marshal(movie)
		if err != nil {
			return fmt.Errorf("Failed to encode payload: %v", err)
		}

		_, err = c.httpClient.Post(
			fmt.Sprintf("http://%s/%s/%s", c.url, movieIndex, movieMapping),
			"application/json",
			bytes.NewBuffer(b),
		)
		if err != nil {
			return fmt.Errorf("Failed sending request: %v", err)
		}
	}

	return nil
}

func (c *ElasticsearchClient) GetAllMovies() ([]*proto.MoviesList_Movie, error) {
	query := "_search/?size=1000"

	resp, err := c.httpClient.Get(
		fmt.Sprintf("http://%s/%s/%s", c.url, movieIndex, query),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed sending request: %v", err)
	}

	var MovieSearchResults struct {
		Hits struct {
			Hits []struct {
				Movie *proto.MoviesList_Movie `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	err = json.NewDecoder(resp.Body).Decode(&MovieSearchResults)
	if err != nil {
		return nil, fmt.Errorf("Failed json decoding search results: %v", err)
	}

	var movies []*proto.MoviesList_Movie

	for _, movieHit := range MovieSearchResults.Hits.Hits {
		movies = append(movies, movieHit.Movie)
	}

	return movies, nil
}
