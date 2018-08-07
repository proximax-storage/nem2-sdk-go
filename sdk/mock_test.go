// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	"time"
)

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
		serv = &mockService{}
		//ignore error
		serv.setupMockServer()
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

func (m *mockService) setupMockServer() error {
	m.mux = http.NewServeMux()

	server := httptest.NewServer(m.mux)

	conf, err := LoadTestnetConfig(server.URL)
	if err != nil {
		return err
	}

	serv.Client = NewClient(nil, conf)
	time.AfterFunc(time.Minute*5, server.Close)

	return nil
}
