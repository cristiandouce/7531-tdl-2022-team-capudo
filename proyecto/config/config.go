package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	PORT string
}

var defaults = map[string]string{
	"PORT": "8000"}

// Package init() functions are guaranteed to be called only once and all called
// from a single thread ( they're thread-safe unless you make them multi-threaded).
// But that makes you dependent on boot order. And you should not write codes in an
// init ( ) that you need a guarantee of execution at any given time
func init() {
	fmt.Println("config invoked for the first time", defaults)
}

func Get(key string) string {
	if value, hasKey := os.LookupEnv(key); hasKey {
		return value
	}

	return defaults[key]
}
