package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahmedshaaban/home24/services"
)

type mockScrapperService struct {
}

type mockHTTPClient struct {
}

func (m *mockScrapperService) Scrap(r io.Reader) (*services.ScrapperResult, error) {
	return &services.ScrapperResult{PageTitle: "test"}, nil
}

func (m *mockHTTPClient) Get(url string) (resp *http.Response, err error) {
	r := ioutil.NopCloser(strings.NewReader(""))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestScrapperHanlder_Scrap(t *testing.T) {
	sHander := NewScrapperHandler(&mockScrapperService{}, &mockHTTPClient{})
	req, err := http.NewRequest("GET", "/scrap?url=https://www.google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sHander.Scrap)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"page_title":"test","headings_count":null,"internal_external_links_count":0,"inaccessible_links_count":0,"contains_login_form":false}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
