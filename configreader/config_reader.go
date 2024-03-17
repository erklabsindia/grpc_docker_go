// Copyright (c) 2022 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package configreader

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// These global variables makes it easy
// to mock these dependencies
// in unit tests.
var (
	godotenvLoad     = godotenv.Load
	envconfigProcess = envconfig.Process
)

// GoDotEnv is an interface that defines
// the functions we use from godotenv package.
// It enables mocking this dependency in unit testing.
type GoDotEnv interface {
	Load(filenames ...string) (err error)
}

// EnvConfig is an interface that defines
// the functions we use from envconfig package.
// It enables mocking this dependency in unit testing.
type EnvConfig interface {
	Process(prefix string, spec interface{}) error
}

// pageSize=20&q=bitcoin&apiKey=
// Config holds configuration data.
type Config struct {
	NewsBaseUrl     string `envconfig:"NEWS_BASE_URL" required:"true"`
	NewsApiKey      string `envconfig:"NEWS_API_KEY" required:"true"`
	NewsHttpTimeout int    `envconfig:"NEWS_HTTP_TIMEOUT" required:"true"`
	SQL_DB          string `envconfig:"SQL_DB" required:"true"`
	SQL_HOST        string `envconfig:"SQL_HOST" required:"true"`
	SQL_PORT        int    `envconfig:"SQL_PORT" required:"true"`
	SQL_USERNAME    string `envconfig:"SQL_USERNAME" required:"true"`
	SQL_PASS        string `envconfig:"SQL_PASS" required:"true"`
}

// ReadEnv reads envionment variables into Config struct.
func ReadEnv() (*Config, error) {
	err := godotenvLoad("configreader/config.env")
	if err != nil {
		return nil, errors.Wrap(err, "reading .env file")
	}
	var config Config
	err = envconfigProcess("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "processing env vars")
	}
	return &config, nil
}
