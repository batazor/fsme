// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SendEventFSMHandlerFunc turns a function with the right signature into a send event f s m handler
type SendEventFSMHandlerFunc func(SendEventFSMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendEventFSMHandlerFunc) Handle(params SendEventFSMParams) middleware.Responder {
	return fn(params)
}

// SendEventFSMHandler interface for that can handle valid send event f s m params
type SendEventFSMHandler interface {
	Handle(SendEventFSMParams) middleware.Responder
}

// NewSendEventFSM creates a new http.Handler for the send event f s m operation
func NewSendEventFSM(ctx *middleware.Context, handler SendEventFSMHandler) *SendEventFSM {
	return &SendEventFSM{Context: ctx, Handler: handler}
}

/*SendEventFSM swagger:route POST /{_id}/event event sendEventFSM

Send event to fsm by ID

*/
type SendEventFSM struct {
	Context *middleware.Context
	Handler SendEventFSMHandler
}

func (o *SendEventFSM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendEventFSMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
