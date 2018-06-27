package retriever

import (
	"fmt"
	"io"
	"net/http"
)

type HttpRetriever struct {
	URL    string
	client *http.Client
}

// HttpClient updates client http.Client if the given http.Client is valid
func HttpClient(client *http.Client) func(*HttpRetriever) error {
	return func(h *HttpRetriever) error {
		if client != nil {
			h.client = client
		}
		return nil
	}
}

func NewHttpRetriever(url string, options ...func(*HttpRetriever) error) (Retriever, error) {

	retriever := &HttpRetriever{URL: url}

	for _, op := range options {
		if err := op(retriever); err != nil {
			return nil, err
		}
	}

	return retriever, nil
}

func (h *HttpRetriever) Retrieve() (io.Reader, error) {
	resp, err := h.client.Get(h.URL)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed - status code %d)", resp.StatusCode)
	}

	return resp.Body, nil
}

