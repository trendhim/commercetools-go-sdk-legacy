package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/trendhim/commercetools-go-sdk-legacy/commercetools"
	"golang.org/x/oauth2"
)

// RequestData holds test HTTP request data
type RequestData struct {
	URL    url.URL
	Body   []byte
	Method string
	JSON   string
}

// HTTPHandler type defines callback from doing a mock HTTP request
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// MockClient returns a client to mock HTTP requests with a callback function
func MockClient(
	t *testing.T,
	fixture string,
	output *RequestData,
	callback HTTPHandler) (*commercetools.Client, *httptest.Server) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		if output != nil {
			output.Method = r.Method
			output.URL = *r.URL
		}

		if r.Method == "POST" || r.Method == "PATCH" {

			// Check if the body is valid JSON
			var dummy map[string]interface{}
			if err := json.Unmarshal(body, &dummy); err != nil {
				log.Printf("Error on unmarshal: %v\n", body)
			}

			if output != nil {
				output.JSON = string(body)
			}
		}

		if callback != nil {
			callback(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture)
		}

	}

	ts := httptest.NewServer(http.HandlerFunc(handler))

	httpClient := oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: "unittest",
	}),
	)

	client := commercetools.New(&commercetools.Config{
		ProjectKey: "unittest",
		URL:        ts.URL,
		HTTPClient: httpClient,
	})

	return client, ts
}
