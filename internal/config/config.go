// Package config loads env-driven configuration for tv-stream.
package config

import "os"

type Config struct {
	Addr    string
	Product string
}

func FromEnv() Config {
	c := Config{Addr: ":8080", Product: "tv"}
	if v := os.Getenv("ADDR"); v != "" {
		c.Addr = v
	}
	return c
}
