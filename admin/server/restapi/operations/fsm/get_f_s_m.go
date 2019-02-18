// Code generated by go-swagger; DO NOT EDIT.

package fsm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFSMHandlerFunc turns a function with the right signature into a get f s m handler
type GetFSMHandlerFunc func(GetFSMParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFSMHandlerFunc) Handle(params GetFSMParams) middleware.Responder {
	return fn(params)
}

// GetFSMHandler interface for that can handle valid get f s m params
type GetFSMHandler interface {
	Handle(GetFSMParams) middleware.Responder
}

// NewGetFSM creates a new http.Handler for the get f s m operation
func NewGetFSM(ctx *middleware.Context, handler GetFSMHandler) *GetFSM {
	return &GetFSM{Context: ctx, Handler: handler}
}

/*GetFSM swagger:route GET /{id} fsm getFSM

GetFSM get f s m API

*/
type GetFSM struct {
	Context *middleware.Context
	Handler GetFSMHandler
}

func (o *GetFSM) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFSMParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
