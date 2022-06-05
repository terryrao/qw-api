package qwapi

import "net/http"

type Option struct {
	logger     Logger
	httpClient *http.Client
}

type APIOptions func(*Option)

func WithLogger(l Logger) APIOptions {
	return func(o *Option) {
		o.logger = l
	}
}

func WithHttpClient(h *http.Client) APIOptions {
	return func(o *Option) {
		o.httpClient = h
	}
}
