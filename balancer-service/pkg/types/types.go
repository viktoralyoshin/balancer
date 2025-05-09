package types

import (
	"net/url"
)

type Config struct {
	Servers []string `json:"servers"`
}

type Server struct {
	URL       *url.URL
	Available int32
}
