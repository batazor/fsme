// Code generated by go-swagger; DO NOT EDIT.

package fsm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DestroyFSMHandlerFunc turns a function with the right signature into a destroy f s m handler
type DestroyFSMHandlerFunc func(DestroyFSMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DestroyFSMHandlerFunc) Handle(params DestroyFSMParams) middleware.Responder {
	return fn(params)
}

// DestroyFSMHandler interface for that can handle valid destroy f s m params
type DestroyFSMHandler interface {
	Handle(DestroyFSMParams) middleware.Responder
}

// NewDestroyFSM creates a new http.Handler for the destroy f s m operation
func NewDestroyFSM(ctx *middleware.Context, handler DestroyFSMHandler) *DestroyFSM {
	return &DestroyFSM{Context: ctx, Handler: handler}
}

/*DestroyFSM swagger:route DELETE /{id} fsm destroyFSM

Delete a fsm

*/
type DestroyFSM struct {
	Context *middleware.Context
	Handler DestroyFSMHandler
}

func (o *DestroyFSM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDestroyFSMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
