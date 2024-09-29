package main

import (
	"os"
)

func commandExit(cfg *Config, name string) error {
	os.Exit(0)
	return nil
}
