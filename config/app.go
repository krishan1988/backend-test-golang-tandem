// Package config is responsible for loading all necessary configurations
// to the backend-test-golang service.
package config

// App contains all configurations.
type App struct {
	Server Server   `json:"server"`
	DB     Database `json:"db"`
}
