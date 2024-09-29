package main

import (
	"fmt"
)

func commandMapb(cfg *Config, name string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you are on the first page")
	}
	data, err := cfg.Client.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}
	for _, v := range data.Locations {
		fmt.Println(v.Name)
	}
	cfg.Next = &data.Next
	cfg.Previous = &data.Previous
	return nil
}
