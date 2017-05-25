package main

import (
	"encoding/json"
	"os"
)

// Config for server
type Config struct {
	Listen    string `json:"listen"`
	Target    string `json:"target"`
	NoComp    bool   `json:"nocomp"`
	SockBuf   int    `json:"sockbuf"`
	KeepAlive int    `json:"keepalive"`
	Log       string `json:"log"`
	Pprof     bool   `json:"pprof"`
}

func parseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}
