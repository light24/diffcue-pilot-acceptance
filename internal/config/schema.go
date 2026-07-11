package config

import "time"

type Settings struct {
	Mode string
	Endpoint string
	Timeout time.Duration
	Strict bool
	MaxRetries int // bounded attempts after the initial request
}

func Default() Settings { return Settings{Mode: "safe", Endpoint: "/v1/events", Timeout: 5 * time.Second, Strict: true, MaxRetries: 3} }
