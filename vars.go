package qwapi

import "net/http"

var (
	defaultHttpClient = &http.Client{
		Timeout: defaultHttpTimeout,
	}
)
