package http

import "net/http"

// Default constants
const (
	StatusOK                  int    = http.StatusOK
	StatusNotFound            int    = http.StatusNotFound
	StatusMethodNotAllowed    int    = http.StatusMethodNotAllowed
	StatusInternalServerError int    = http.StatusInternalServerError
	MethodGet                        = http.MethodGet
	MethodHead                       = http.MethodHead
	FQDNPrivate               string = "http://localhost/"
	FQDNPublic                string = "https://dives.dev/"
)

// Default variables
var (
	MessageNotFound            = http.StatusText(http.StatusNotFound)
	MessageOK                  = http.StatusText(http.StatusOK)
	MessageMethodNotAllowed    = http.StatusText(http.StatusMethodNotAllowed)
	MessageInternalServerError = http.StatusText(http.StatusInternalServerError)
)
