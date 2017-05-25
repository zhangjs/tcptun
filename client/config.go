package main

import (
	"encoding/json"
	"os"
)

// Config for client
type Config struct {
	LocalAddr   string `json:"localaddr"`
	RemoteAddr  string `json:"remoteaddr"`
	Conn        int    `json:"conn"`
	AutoExpire  int    `json:"autoexpire"`
	ScavengeTTL int    `json:"scavengettl"`
	NoComp      bool   `json:"nocomp"`
	SockBuf     int    `json:"sockbuf"`
	KeepAlive   int    `json:"keepalive"`
	Log         string `json:"log"`
}

func parseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}
