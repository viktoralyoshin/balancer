package config

import (
	"balancer/pkg/types"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
)

func LoadConfig(configPath string) (*types.Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg types.Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseServers(serverURLs []string) ([]*types.Server, error) {
	servers := make([]*types.Server, 0, len(serverURLs))

	for _, urlStr := range serverURLs {
		u, err := url.Parse(urlStr)
		if err != nil {
			log.Printf("Failed to parse URL %s: %v", urlStr, err)
			continue
		}
		servers = append(servers, &types.Server{URL: u, Available: 1})
	}

	if len(servers) == 0 {
		return nil, errors.New("no valid servers in config")
	}

	return servers, nil
}
