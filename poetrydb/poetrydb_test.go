// Copyright (c) 2022 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package poetrydb

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	testCases := []struct {
		name              string
		mockClosure       func(m *mockedHttpClient)
		mockIoUtilReadAll func(r io.Reader) ([]byte, error)
		mockJsonMarshal   func(v interface{}) ([]byte, error)
		mockJsonUnmarshal func(data []byte, v interface{}) error
		expectedResponse  PoetriesResponse
		expectedError     error
	}{
		{
			name: "happy path",
			mockClosure: func(m *mockedHttpClient) {
				body := `[{"title":"Poetry title","author":"Author","lines":["line 1","line 2"],"linecount":"2"}, {"title":"Poetry title 2","author":"Author","lines":["line 1","line 2"],"linecount":"2"}]`
				m.Response = &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(body))),
				}
			},
			expectedResponse: PoetriesResponse{
				List: []PoetryResponse{
					{
						Title:  "Poetry title",
						Author: "Author",
						Lines: []string{
							"line 1",
							"line 2",
						},
						LineCount: "2",
					},
					{
						Title:  "Poetry title 2",
						Author: "Author",
						Lines: []string{
							"line 1",
							"line 2",
						},
						LineCount: "2",
					},
				},
			},
		},
		{
			name: "error making request",
			mockClosure: func(m *mockedHttpClient) {
				m.Error = errors.New("random error")
			},
			expectedError:    errors.New(`performing request to "url/random/2": random error`),
			expectedResponse: PoetriesResponse{},
		},
		{
			name: "http status code different than 2xx",
			mockClosure: func(m *mockedHttpClient) {
				m.Response = &http.Response{
					StatusCode: 500,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(nil))),
				}
			},
			expectedError:    errors.New("got status code 500"),
			expectedResponse: PoetriesResponse{},
		},
		{
			name: "error reading response",
			mockClosure: func(m *mockedHttpClient) {
				m.Response = &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(nil))),
				}
			},
			mockIoUtilReadAll: func(r io.Reader) ([]byte, error) {
				return nil, errors.New("random error")
			},
			expectedError:    errors.New("reading response: random error"),
			expectedResponse: PoetriesResponse{},
		},
		{
			name: "error parsing response",
			mockClosure: func(m *mockedHttpClient) {
				m.Response = &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(nil))),
				}
			},
			mockJsonUnmarshal: func(data []byte, v interface{}) error {
				return errors.New("random error")
			},
			expectedError:    errors.New("parsing response: random error"),
			expectedResponse: PoetriesResponse{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockIoUtilReadAll != nil {
				ioUtilReadAll = tc.mockIoUtilReadAll
			} else {
				ioUtilReadAll = ioutil.ReadAll
			}
			if tc.mockJsonUnmarshal != nil {
				jsonUnmarshal = tc.mockJsonUnmarshal
			} else {
				jsonUnmarshal = json.Unmarshal
			}
			mockedHttpClient := new(mockedHttpClient)
			if tc.mockClosure != nil {
				tc.mockClosure(mockedHttpClient)
			}
			poetryDb := &poetryDb{
				baseUrl:    "url",
				httpClient: mockedHttpClient,
			}
			poetryResponse, err := poetryDb.Random(2)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no errors, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatal("expected error to occur")
				}
			}
			require.Equal(t, tc.expectedResponse, poetryResponse)
		})
	}
}

type mockedHttpClient struct {
	Response *http.Response
	Error    error
}

func (m *mockedHttpClient) Get(url string) (resp *http.Response, err error) {
	return m.Response, m.Error
}
