package main

import (
	"fmt"
)

func commandMap(cfg *Config, name string) error {
	data, err := cfg.Client.GetLocations(cfg.Next)
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
