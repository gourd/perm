package perm

import (
	"net/http"
)

// HandlerFunc is the basic unit for permission management
// It takes
type HandlerFunc func(r *http.Request, perm string, info ...interface{}) error

// Allow calls f(r, pern, info...).
func (h HandlerFunc) Allow(r *http.Request, perm string, info ...interface{}) error {
	return h(r, perm, info...)
}
