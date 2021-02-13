package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ahmedshaaban/home24/services"
)

type scrapperService interface {
	Scrap(r io.Reader) (*services.ScrapperResult, error)
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// ScrapperHanlder ...
type ScrapperHanlder struct {
	scrapperService scrapperService
	httpClient      httpClient
}

// NewScrapperHandler ...
func NewScrapperHandler(sService scrapperService, httpClient httpClient) *ScrapperHanlder {
	return &ScrapperHanlder{
		scrapperService: sService,
		httpClient:      httpClient,
	}
}

// Scrap ...
func (s *ScrapperHanlder) Scrap(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	response, err := s.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	result, err := s.scrapperService.Scrap(response.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
