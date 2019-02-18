// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/batazor/fsme/admin/server/restapi/operations"
	"github.com/batazor/fsme/admin/server/restapi/operations/fsm"
)

//go:generate swagger generate server --target ../../server --name FsmServer --spec ../docs/swagger.yaml

func configureFlags(api *operations.FsmServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FsmServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FsmAddFSMHandler = fsm.AddFSMHandlerFunc(func(params fsm.AddFSMParams) middleware.Responder {
		return middleware.NotImplemented("operation fsm.AddFSM has not yet been implemented")
	})
	api.FsmDestroyFSMHandler = fsm.DestroyFSMHandlerFunc(func(params fsm.DestroyFSMParams) middleware.Responder {
		return middleware.NotImplemented("operation fsm.DestroyFSM has not yet been implemented")
	})
	api.FsmGetFSMHandler = fsm.GetFSMHandlerFunc(func(params fsm.GetFSMParams) middleware.Responder {
		return middleware.NotImplemented("operation fsm.GetFSM has not yet been implemented")
	})
	api.FsmGetFSMListHandler = fsm.GetFSMListHandlerFunc(func(params fsm.GetFSMListParams) middleware.Responder {
		return middleware.NotImplemented("operation fsm.GetFSMList has not yet been implemented")
	})
	api.FsmUpdateFSMHandler = fsm.UpdateFSMHandlerFunc(func(params fsm.UpdateFSMParams) middleware.Responder {
		return middleware.NotImplemented("operation fsm.UpdateFSM has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
