// Package config is responsible for loading all necessary configurations
// to the backend-test-golang service.
package config

// Database database contains all the configurations related database.
type Database struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}
