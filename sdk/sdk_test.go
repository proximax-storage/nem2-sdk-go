package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	address = "http://10.32.150.136:3000"
)

func setupWithAddress(adr string) *Client {
	conf, err := LoadTestnetConfig(adr)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf)
}

func setup() (*Client, string) {
	conf, err := LoadTestnetConfig(address)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf), address
}

// Create a mock server
func setupMockServer() (client *Client, mux *http.ServeMux, serverURL string, teardown func(), err error) {
	// individual tests will provide API mock responses
	mux = http.NewServeMux()

	server := httptest.NewServer(mux)

	conf, err := LoadTestnetConfig(server.URL)
	if err != nil {
		return nil, nil, "", nil, err
	}

	client = NewClient(nil, conf)

	return client, mux, server.URL, server.Close, nil
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

// Uint64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Uint64(v uint64) *uint64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

type sParam struct {
	desc     string
	req      bool
	Type     string
	defValue interface{}
}

type sRouting struct {
	resp   string
	params map[string]sParam
}

func (r *sRouting) checkParams(req *http.Request) (badParams []string, err error) {
	for key, val := range r.params {

		if key == "body" {
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				err = errors.New("failed during reading Body")
				return badParams, err
			}
			if (len(b) == 0) || (bytes.Contains(b, []byte("null"))) {
				badParams = append(badParams, "body is empty")
			}
			//	todo: add check struct to match the request requirements
		} else if valueParam := req.FormValue(key); (val.req) && (valueParam == "") {
			badParams = append(badParams, key)
		} else if val.Type > "" {
			//	check type is later
			if valueParam != val.Type {
				err = errors.New("bad type param")
				return
			}
		}
	}

	return
}

type mockService struct {
	*Client
	mux  *http.ServeMux
	lock sync.Locker
}

func addRouters(routers map[string]sRouting) {

	if serv == nil {
		serv = NewMockServer()
	}

	for path, route := range routers {
		apiRoute := route
		serv.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			// Mock JSON response
			if params, err := apiRoute.checkParams(r); (len(params) > 0) || (err != nil) {
				w.WriteHeader(http.StatusBadRequest)
				if len(params) > 0 {
					p := strings.Join(params, ",")
					fmt.Fprintf(w, "bad params - %s", p)
				}
				if err != nil {
					fmt.Fprint(w, "error during params validate - ", err)
				}
			} else {
				w.Write([]byte(apiRoute.resp))
			}
		})

	}

}

var (
	serv          *mockService
	ctx           = context.TODO()
	routeNeedBody = map[string]sParam{"body": {desc: "required body"}}
)

func NewMockServer() *mockService {
	client, mux, _, teardown, err := setupMockServer()

	if err != nil {
		panic(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//	mock router as default
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found in mock routers", r.URL)
		fmt.Println(r.URL)
	})
	time.AfterFunc(time.Minute*5, teardown)

	return &mockService{mux: mux, Client: client}
}

func validateResp(resp *http.Response, t *testing.T) bool {
	if resp == nil {
		return false
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp.Body)
		return false
	}
	return true
}
