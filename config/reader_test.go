// Package config_test contains unit test for configuration functionalities.
package config_test

import (
	"errors"
	"testing"

	"github.com/KryptoKnight/backend-test-golang/config"
	"github.com/stretchr/testify/assert"
)

// TestReader_Read test read functionality of the reader
func TestReader_Read(t *testing.T) {
	dummyErr := errors.New("dummy error")
	tests := []struct {
		Desc        string
		LoadFunc    config.FileLoadFunc
		ExpectedApp config.App
		ExpectedErr error
	}{
		{
			Desc: "Check when valid json file parse to the application",
			LoadFunc: func(s string) ([]byte, error) {
				return []byte(
					`
{
    "server": {
        "port": 8083
    },
    "db": {
        "uri": "mongodb://127.0.0.1:27100",
        "name": "sample"
    }
}`), nil
			},
			ExpectedApp: config.App{
				DB: config.Database{
					URI:  "mongodb://127.0.0.1:27100",
					Name: "sample",
				},
				Server: config.Server{
					Port: 8083,
				},
			},
		},
		{
			Desc: "Check error when load function return error",
			LoadFunc: func(s string) ([]byte, error) {
				return []byte(``), dummyErr
			},
			ExpectedApp: config.App{},
			ExpectedErr: dummyErr,
		},
	}

	for _, test := range tests {
		reader := config.NewReader[config.App]("sample")
		reader.WithCustomFileLoadFunc(test.LoadFunc)
		actualApp, actualErr := reader.Read()
		assert.Equal(t, test.ExpectedErr, actualErr, test.Desc)
		if actualApp == nil {
			continue
		}
		assert.Equal(t, test.ExpectedApp, *actualApp, test.Desc)
	}
}
