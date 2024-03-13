// Copyright (c) 2022 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package server

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"testing"

	"bitbucket.org/tiagoharris/docker-grpc-service-tutorial/configreader"
	"bitbucket.org/tiagoharris/docker-grpc-service-tutorial/poetrydb"
	poetry "bitbucket.org/tiagoharris/docker-grpc-service-tutorial/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestNewServer(t *testing.T) {
	testCases := []struct {
		name                    string
		mockNetListen           func(network string, address string) (net.Listener, error)
		mockConfigreaderReadEnv func() (*configreader.Config, error)
		expectedError           error
	}{
		{
			name: "happy path",
			mockNetListen: func(network, address string) (net.Listener, error) {
				return &net.TCPListener{}, nil
			},
			mockConfigreaderReadEnv: func() (*configreader.Config, error) {
				return &configreader.Config{}, nil
			},
		},
		{
			name: "tcp listening error",
			mockNetListen: func(network, address string) (net.Listener, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("tcp listening: random error"),
		},
		{
			name: "reading env vars error",
			mockNetListen: func(network, address string) (net.Listener, error) {
				return &net.TCPListener{}, nil
			},
			mockConfigreaderReadEnv: func() (*configreader.Config, error) {
				return nil, errors.New("random error")
			},
			expectedError: errors.New("reading env vars: random error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			netListen = tc.mockNetListen
			configreaderReadEnv = tc.mockConfigreaderReadEnv
			server, err := NewServer(4040)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no errors to occur, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v" to occur, got nil`, tc.expectedError)
				}
			}
			require.NotNil(t, server)
		})
	}
}

func TestRandomPoetries(t *testing.T) {
	testCases := []struct {
		name                   string
		mockClosure            func(m *mockPoetryDb)
		mockJsonMarshal        func(v interface{}) ([]byte, error)
		mockProtojsonUnmarshal func(b []byte, m protoreflect.ProtoMessage) error
		expectedPoetryList     *poetry.PoetryList
		expectedError          error
	}{
		{
			name: "happy path",
			mockClosure: func(m *mockPoetryDb) {
				response := poetrydb.PoetriesResponse{
					List: []poetrydb.PoetryResponse{
						{
							Title:     "title",
							Author:    "author",
							Lines:     []string{"line 1", "line 2"},
							LineCount: "2",
						},
					},
				}
				m.Response = response
			},
			expectedPoetryList: &poetry.PoetryList{
				List: []*poetry.Poetry{
					{
						Title:     "title",
						Author:    "author",
						Lines:     []string{"line 1", "line 2"},
						Linecount: int32(2),
					},
				},
			},
		},
		{
			name: "error calling PoetryDb service",
			mockClosure: func(m *mockPoetryDb) {
				m.Response = poetrydb.PoetriesResponse{}
				m.Err = errors.New("random error")
			},
			expectedPoetryList: &poetry.PoetryList{},
			expectedError:      errors.New("requesting random poetry: random error"),
		},
		{
			name: "error marshalling json",
			mockClosure: func(m *mockPoetryDb) {
				response := poetrydb.PoetriesResponse{
					List: []poetrydb.PoetryResponse{
						{
							Title:     "title",
							Author:    "author",
							Lines:     []string{"line 1", "line 2"},
							LineCount: "2",
						},
					},
				}
				m.Response = response
			},
			mockJsonMarshal: func(v interface{}) ([]byte, error) {
				return nil, errors.New("random error")
			},
			expectedPoetryList: &poetry.PoetryList{},
			expectedError:      errors.New("marshalling json: random error"),
		},
		{
			name: "error unmarshalling proto",
			mockClosure: func(m *mockPoetryDb) {
				response := poetrydb.PoetriesResponse{
					List: []poetrydb.PoetryResponse{
						{
							Title:     "title",
							Author:    "author",
							Lines:     []string{"line 1", "line 2"},
							LineCount: "2",
						},
					},
				}
				m.Response = response
			},
			mockProtojsonUnmarshal: func(b []byte, m protoreflect.ProtoMessage) error {
				return errors.New("random error")
			},
			expectedPoetryList: &poetry.PoetryList{},
			expectedError:      errors.New("unmarshalling proto: random error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockJsonMarshal != nil {
				jsonMarshal = tc.mockJsonMarshal
			} else {
				jsonMarshal = json.Marshal
			}
			if tc.mockProtojsonUnmarshal != nil {
				protojsonUnmarshal = tc.mockProtojsonUnmarshal
			} else {
				protojsonUnmarshal = protojson.Unmarshal
			}
			ctx := context.TODO()
			mockPoetryDb := new(mockPoetryDb)
			tc.mockClosure(mockPoetryDb)
			server := &server{
				poetryDb: mockPoetryDb,
			}
			poetryList, err := server.RandomPoetries(ctx, &poetry.RandomPoetriesRequest{NumberOfPoetries: 1})
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no errors to occur, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v" to occur, got nil`, tc.expectedError)
				}
			}
			require.True(t, proto.Equal(tc.expectedPoetryList, poetryList))
		})
	}
}

type mockPoetryDb struct {
	Response poetrydb.PoetriesResponse
	Err      error
}

func (m *mockPoetryDb) Random(number int) (poetrydb.PoetriesResponse, error) {
	return m.Response, m.Err
}
