package slick

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	request  *http.Request
	response http.ResponseWriter
	ctx      context.Context
}

// renders a templ Component with friendly support of slick.Context
func (c *Context) Render(tcp templ.Component) error {
	return tcp.Render(c.ctx, c.response)
}

type ErrorHandler func(*Context, error) error

// carrier for all Slick methods
type Slick struct {
	ErrorHandler ErrorHandler
	Router       *httprouter.Router
}

type Handler func(c *Context) error

// Slick constructor
func New() *Slick {
	return &Slick{
		ErrorHandler: defaultErrorHandler,
		Router:       httprouter.New(),
	}
}

// fires up the web server
func (s *Slick) Start(port string) error {
	return http.ListenAndServe(port, s.Router)
}

// wrapper for httprouter.GET(); 'plugs' is reserved for middleware functions.
func (s *Slick) Get(path string, h Handler, plugs ...Handler) {
	s.Router.GET(path, s.makeHTTPRouterHandle(h))
}

// configures and returns a ready-to-go httprouter.Handle function
// for a custom Endpoint Handler 'h' for httprouter to execute
func (s *Slick) makeHTTPRouterHandle(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// this will be great for creating a context pool!
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := h(ctx); err != nil {
			// TODO: handle the error returned by the ErrorHandler
			s.ErrorHandler(ctx, err)
		}
	}
}

func defaultErrorHandler(ctx *Context, err error) error {
	slog.Info("ERROR", "err", err)
	return nil
}

// vim: ts=4 sw=4 fdm=indent
