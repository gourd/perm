package perm

import (
	"errors"
	"github.com/gorilla/context"
	"net/http"
)

var HandlerNotFound error
var contextKey *int

func init() {
	i := 0
	contextKey = &i
	HandlerNotFound = errors.New("Permission handler not found")
}

// Mux is the primary permission interface
// for user to obtain permission
type Mux interface {
	HandleFunc(perm string, h HandlerFunc)
	Handle(perm string, h Handler)
	Allow(r *http.Request, perm string, info ...interface{}) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// NewMux returns a new DefaultMux
func NewMux() Mux {
	var dh HandlerFunc = func(r *http.Request, perm string, info ...interface{}) error {
		return HandlerNotFound
	}
	return &DefaultMux{
		handlers: make(map[string]Handler),
		defaultH: dh,
	}
}

// DefaultMux route permission request to different
// permission Handler. The design mimics ServeMux
// in the core http pacakge
type DefaultMux struct {
	// unexported variables
	handlers map[string]Handler
	defaultH Handler
}

// Handle register a Handler to the DefaultMux.
// This Handler will be provided when calling ServePerm with
// the perm string equals perm.
// The design mimics *ServeMux.Handle
func (p *DefaultMux) Handle(perm string, h Handler) {
	// overwrite by default (until I figure something else)
	p.handlers[perm] = h
}

// HandleFunc register a HandlerFunc to the DefaultMux.
// This Handler will be provided when calling ServePerm with
// the perm string equals perm.
// The design mimics *ServeMux.HandleFunc
func (p *DefaultMux) HandleFunc(perm string, h HandlerFunc) {
	// overwrite by default (until I figure something else)
	p.handlers[perm] = h
}

// Allow dispatches the permission request to the registered
// handlers whose perm string matches / most close to the registered
// Handler
func (p *DefaultMux) Allow(r *http.Request, perm string, info ...interface{}) error {
	if handler, ok := p.handlers[perm]; ok {
		return handler.Allow(r, perm, info...)
	}
	// TODO: find relevant permission string by pattern (i.e. `*`)

	return p.defaultH.Allow(r, perm, info...)
}

// ServeHTTP provide itself to the context
func (p *DefaultMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.Set(r, contextKey, p)
}

// GetMuxOk
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
