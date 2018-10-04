// Package sdk provides a client library for the Catapult REST API.
package sdk

import (
	"bytes"
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/json-iterator/go"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Provides service configuration
type Config struct {
	BaseURL *url.URL
	NetworkType
}

// Mainnet config default
func LoadMainnetConfig(baseUrl string) (*Config, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	c := &Config{BaseURL: u, NetworkType: MainNet}

	return c, nil
}

// Testnet config default
func LoadTestnetConfig(baseUrl string) (*Config, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	c := &Config{BaseURL: u, NetworkType: TestNet}
	return c, nil
}

// Catapult API Client configuration
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.
	config *Config
	common service // Reuse a single struct instead of allocating one for each service on the heap.
	// Services for communicating to the Catapult REST APIs
	Blockchain  *BlockchainService
	Mosaic      *MosaicService
	Namespace   *NamespaceService
	Network     *NetworkService
	Transaction *TransactionService
	Account     *AccountService
}

type service struct {
	client *Client
}

// NewClient returns a new Catapult API client.
// If httpClient is nil then it will create http.DefaultClient
func NewClient(httpClient *http.Client, conf *Config) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient, config: conf}
	c.common.client = c
	c.Blockchain = (*BlockchainService)(&c.common)
	c.Mosaic = (*MosaicService)(&c.common)
	c.Namespace = (*NamespaceService)(&c.common)
	c.Network = (*NetworkService)(&c.common)
	c.Transaction = (*TransactionService)(&c.common)
	c.Account = (*AccountService)(&c.common)

	return c
}

// DoNewRequest creates new request, Do it & return result in V
func (s *Client) DoNewRequest(ctx context.Context, method string, path string, body interface{}, v interface{}) (*http.Response, error) {
	req, err := s.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.Do(ctx, req, v)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Do sends an API Request and returns a parsed response
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {

	// set the Context for this request
	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 226 || resp.StatusCode < 200 {
		return resp, errors.New(resp.Status)
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(&v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	u, err := c.config.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
