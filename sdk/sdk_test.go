package sdk

import (
	"net/http"
	"net/http/httptest"
)

const (
	address = "http://10.32.150.136:3000"
)

func setup() (*Client, string) {

	conf, err := LoadTestnetConfig(address)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf), address
}

// Create a mock server
func setupMockServer() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// individual tests will provide API mock responses
	mux = http.NewServeMux()

	server := httptest.NewServer(mux)

	conf, err := LoadTestnetConfig(server.URL)
	if err != nil {
		panic(err)
	}

	client = NewClient(nil, conf)

	return client, mux, server.URL, server.Close
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
