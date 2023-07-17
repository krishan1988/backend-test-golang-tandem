// Package config is responsible for loading all necessary configurations
// to the backend-test-golang service.
package config

// Server contains all preload configurations related backend-test-golang server.
type Server struct {
	PublicKeyPath string `yaml:"publicKeyPath"`
	Port          int    `yaml:"port"`
}
