package perm

import (
	"net/http"
)

// Handler handle permission requests
type Handler interface {

	// Allow returns nil if permission is granted
	// or return an error if permission is denied
	Allow(r *http.Request, perm string, info ...interface{}) error
}
