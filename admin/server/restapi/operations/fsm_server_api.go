// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/batazor/fsme/admin/server/restapi/operations/event"
	"github.com/batazor/fsme/admin/server/restapi/operations/fsm"
)

// NewFsmServerAPI creates a new FsmServer instance
func NewFsmServerAPI(spec *loads.Document) *FsmServerAPI {
	return &FsmServerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		FsmAddFSMHandler: fsm.AddFSMHandlerFunc(func(params fsm.AddFSMParams) middleware.Responder {
			return middleware.NotImplemented("operation FsmAddFSM has not yet been implemented")
		}),
		FsmDestroyFSMHandler: fsm.DestroyFSMHandlerFunc(func(params fsm.DestroyFSMParams) middleware.Responder {
			return middleware.NotImplemented("operation FsmDestroyFSM has not yet been implemented")
		}),
		FsmGetFSMHandler: fsm.GetFSMHandlerFunc(func(params fsm.GetFSMParams) middleware.Responder {
			return middleware.NotImplemented("operation FsmGetFSM has not yet been implemented")
		}),
		FsmGetFSMListHandler: fsm.GetFSMListHandlerFunc(func(params fsm.GetFSMListParams) middleware.Responder {
			return middleware.NotImplemented("operation FsmGetFSMList has not yet been implemented")
		}),
		EventSendEventFSMHandler: event.SendEventFSMHandlerFunc(func(params event.SendEventFSMParams) middleware.Responder {
			return middleware.NotImplemented("operation EventSendEventFSM has not yet been implemented")
		}),
		FsmUpdateFSMHandler: fsm.UpdateFSMHandlerFunc(func(params fsm.UpdateFSMParams) middleware.Responder {
			return middleware.NotImplemented("operation FsmUpdateFSM has not yet been implemented")
		}),
	}
}

/*FsmServerAPI FSME server for UI */
type FsmServerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/io.goswagger.examples.todo-list.v1+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/io.goswagger.examples.todo-list.v1+json" mime type
	JSONProducer runtime.Producer

	// FsmAddFSMHandler sets the operation handler for the add f s m operation
	FsmAddFSMHandler fsm.AddFSMHandler
	// FsmDestroyFSMHandler sets the operation handler for the destroy f s m operation
	FsmDestroyFSMHandler fsm.DestroyFSMHandler
	// FsmGetFSMHandler sets the operation handler for the get f s m operation
	FsmGetFSMHandler fsm.GetFSMHandler
	// FsmGetFSMListHandler sets the operation handler for the get f s m list operation
	FsmGetFSMListHandler fsm.GetFSMListHandler
	// EventSendEventFSMHandler sets the operation handler for the send event f s m operation
	EventSendEventFSMHandler event.SendEventFSMHandler
	// FsmUpdateFSMHandler sets the operation handler for the update f s m operation
	FsmUpdateFSMHandler fsm.UpdateFSMHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *FsmServerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FsmServerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FsmServerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FsmServerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FsmServerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FsmServerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FsmServerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FsmServerAPI
func (o *FsmServerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.FsmAddFSMHandler == nil {
		unregistered = append(unregistered, "fsm.AddFSMHandler")
	}

	if o.FsmDestroyFSMHandler == nil {
		unregistered = append(unregistered, "fsm.DestroyFSMHandler")
	}

	if o.FsmGetFSMHandler == nil {
		unregistered = append(unregistered, "fsm.GetFSMHandler")
	}

	if o.FsmGetFSMListHandler == nil {
		unregistered = append(unregistered, "fsm.GetFSMListHandler")
	}

	if o.EventSendEventFSMHandler == nil {
		unregistered = append(unregistered, "event.SendEventFSMHandler")
	}

	if o.FsmUpdateFSMHandler == nil {
		unregistered = append(unregistered, "fsm.UpdateFSMHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FsmServerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FsmServerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *FsmServerAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *FsmServerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.examples.todo-list.v1+json":
			result["application/io.goswagger.examples.todo-list.v1+json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *FsmServerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.examples.todo-list.v1+json":
			result["application/io.goswagger.examples.todo-list.v1+json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *FsmServerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the fsm server API
func (o *FsmServerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FsmServerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = fsm.NewAddFSM(o.context, o.FsmAddFSMHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{id}"] = fsm.NewDestroyFSM(o.context, o.FsmDestroyFSMHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/{id}"] = fsm.NewGetFSM(o.context, o.FsmGetFSMHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = fsm.NewGetFSMList(o.context, o.FsmGetFSMListHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/{id}/event"] = event.NewSendEventFSM(o.context, o.EventSendEventFSMHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/{id}"] = fsm.NewUpdateFSM(o.context, o.FsmUpdateFSMHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FsmServerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *FsmServerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *FsmServerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *FsmServerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
