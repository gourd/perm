package perm

import (
	"github.com/gorilla/context"
	"net/http"
)

// GetMuxOk retrieve permission from current gorilla context
// and a boolean flag. If not found, it returns flase flag.
func GetMuxOk(r *http.Request) (m Mux, ok bool) {

	// try to get current key
	cm, ok := context.GetOk(r, contextKey)
	if !ok {
		return
	}

	m, ok = cm.(*DefaultMux)
	return
}

// GetMux retrieve permission mux from current gorilla context
func GetMux(r *http.Request) (m Mux) {
	m, _ = GetMuxOk(r)
	return
}
