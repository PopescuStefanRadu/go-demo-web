package config

import "faceit/pkg/api/http"

type Parsed struct {
	Environemnt  string
	ServerConfig http.ServerConfig
}
