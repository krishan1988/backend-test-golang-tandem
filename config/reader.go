// Package config is responsible for loading all necessary configurations
// to the backend-test-golang service.
package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// FileLoadFunc a type to define file loading functionality.
type FileLoadFunc func(string) ([]byte, error)

// FileReadFunc a type to define file reading functionality.
type FileReadFunc[T any] func([]byte) (*T, error)

// Reader reader configurations from the defined
// files in the environment configuration.
type Reader[T any] struct {
	filename string
	cfg      *T
	load     FileLoadFunc
	read     FileReadFunc[T]
}

// NewReader creates an new instance of a configuration reader.
func NewReader[T any](filename string) *Reader[T] {
	r := &Reader[T]{
		filename: filename,
		cfg:      new(T),
	}
	r.setDefaultFileLoader()
	r.setDefaultFileReader()
	return r
}

// WithCustomFileLoadFunc set custom file loader
func (r *Reader[T]) WithCustomFileLoadFunc(loadFunc FileLoadFunc) {
	r.load = loadFunc
}

// WithCustomFileLoader set custom file loader
func (r *Reader[T]) WithCustomFileReadFunc(readFunc FileReadFunc[T]) {
	r.read = readFunc
}

// Read reads the given configuration file.
func (r *Reader[T]) Read() (*T, error) {
	file, err := r.load(r.filename)
	if err != nil {
		return nil, err
	}
	if len(file) == 0 {
		return nil, errors.New("given file is empty")
	}
	return r.read(file)
}

// setDefaultFileLoader set default file loader
func (r *Reader[t]) setDefaultFileLoader() {
	r.load = func(filename string) ([]byte, error) {
		file, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		return file, err
	}
}

// setDefaultFileReader set default file reader ( JSON implementation).
func (r *Reader[t]) setDefaultFileReader() {
	r.read = func(file []byte) (*t, error) {
		t := new(t)
		err := yaml.Unmarshal(file, t)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}
