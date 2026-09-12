package core

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Mode    string `json:"mode"`

	Modules struct {
		Network bool `json:"network"`
		DNS     bool `json:"dns"`
		System  bool `json:"system"`
		Tools   bool `json:"tools"`
		Lab     bool `json:"lab"`
		CTF     bool `json:"ctf"`
		Reports bool `json:"reports"`
	} `json:"modules"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}