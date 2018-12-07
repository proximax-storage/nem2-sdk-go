// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	address = "http://10.32.150.136:3000"
)

var (
	ctx        = context.TODO()
	mockServer = newMock(5 * time.Minute)
)

func setupWithAddress(adr string) *Client {
	conf, err := NewConfig(adr, TestNet)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf)
}

func setup() (*Client, string) {
	conf, err := NewConfig(address, TestNet)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf), address
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

//using different numbers from original javs sdk because of signed and unsigned transformation
//ex. uint64(-8884663987180930485) = 9562080086528621131
func TestBigIntegerToHex_bigIntegerNEMAndXEMToHex(t *testing.T) {
	testBigInt(t, "9562080086528621131", "84b3552d375ffa4b")
	testBigInt(t, "15358872602548358953", "d525ad41d95fcf29")
}

func testBigInt(t *testing.T, str, hexStr string) {
	i, ok := (&big.Int{}).SetString(str, 10)
	assert.True(t, ok)
	result := BigIntegerToHex(i)
	assert.Equal(t, hexStr, result)
}

type mock struct {
	server *httptest.Server
	mux    *http.ServeMux
	lock   sync.Mutex
}

type router struct {
	path              string
	respHttpCode      int
	respBody          string
	reqJsonBodyStruct interface{}
	formParams        []formParam
}

type formParam struct {
	name       string
	isRequired bool
}

func newMock(closeAfter time.Duration) *mock {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//	mock router as default
		resp.WriteHeader(http.StatusNotFound)

		writeStringToResp(resp, fmt.Sprintf("%s not found in mock routers", req.URL))
	})

	if closeAfter != 0 {
		time.AfterFunc(closeAfter, server.Close)
	}

	return &mock{
		mux:    mux,
		server: server,
	}
}

func newMockWithRoute(router *router) *mock {
	mockServer := newMock(0)

	mockServer.addRouter(router)

	return mockServer
}

func (m *mock) close() {
	m.server.Close()
}

func (m *mock) getClientByNetworkType(networkType NetworkType) (*Client, error) {
	conf, err := NewConfig(m.server.URL, networkType)

	if err != nil {
		return nil, err
	}

	client := NewClient(nil, conf)

	return client, nil
}

func (m *mock) getTestNetClient() (*Client, error) {
	return m.getClientByNetworkType(TestNet)
}

func (m *mock) getTestNetClientUnsafe() *Client {
	client, _ := m.getTestNetClient()

	return client
}

func (m *mock) addHandler(path string, handler func(resp http.ResponseWriter, req *http.Request)) {
	m.mux.HandleFunc(path, handler)
}

func (m *mock) addRouter(routers ...*router) {
	if len(routers) == 0 {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	for _, router := range routers {
		if router == nil {
			continue
		}

		m.addHandler(
			router.path,
			func(resp http.ResponseWriter, req *http.Request) {
				// Checking json body
				if router.reqJsonBodyStruct != nil {
					bodyBytes, err := ioutil.ReadAll(req.Body)

					if len(bodyBytes) == 0 || err != nil {
						resp.WriteHeader(http.StatusBadRequest)

						return
					}

					jsonStructType := reflect.TypeOf(router.reqJsonBodyStruct)

					newObj := reflect.New(jsonStructType).Elem().Addr()

					err = json.Unmarshal(bodyBytes, newObj.Interface())

					if err != nil {
						resp.WriteHeader(http.StatusBadRequest)

						writeStringToResp(resp, err.Error())

						return
					}
				} else if len(router.formParams) != 0 { // If not json maybe are there form parameters ?
					errors := make([]string, 0, 1)

					for _, param := range router.formParams {
						val := req.Form[param.name]

						if param.isRequired && len(val) == 0 {
							errors = append(errors, fmt.Sprintf("value of %s is blank", param.name))
						}
					}

					if len(errors) != 0 {
						resp.WriteHeader(http.StatusBadRequest)

						writeStringToResp(resp, strings.Join(errors, ", "))

						return
					}
				}

				if router.respHttpCode != 0 {
					resp.WriteHeader(router.respHttpCode)
				}

				if len(router.respBody) != 0 {
					writeStringToResp(resp, router.respBody)
				}
			},
		)
	}
}

func writeStringToResp(resp http.ResponseWriter, str string) {
	n, err := io.WriteString(resp, str)

	if n != len(str) || err != nil {
		fmt.Printf("failed within writing response body [str=%s, wroteCount=%d, err=%v]\n", str, n, err)
	}
}
