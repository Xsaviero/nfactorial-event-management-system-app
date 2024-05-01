package config

import "time"

type Config struct {
	
}

type HTTPServer struct {
	Address string `json:"address" env-default:"4040"`
	Timeout time.Duration `json:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `json:"idle_timeout" env-default:"60s"`
	User string `json:"user" env-required:"true"`
	Password string `json:"password" env-required:"true" env`
}

func MustLoad() *Config {
	return &Config{}
}