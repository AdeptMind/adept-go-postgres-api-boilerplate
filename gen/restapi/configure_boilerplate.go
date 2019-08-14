// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/config"
	db "github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/handler"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/logger"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/users"
)

//go:generate swagger generate server --target ../../gen --name Boilerplate --spec ../../swagger.yml --exclude-main

func configureFlags(api *operations.BoilerplateAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BoilerplateAPI) http.Handler {
	c := config.GetConfig()
	logger.ConfigureLogger(c.LogLevel)

	gorm, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=%v password=%v", c.DbHost, c.DbPort, c.DbUser, c.DbName, c.DbSslMode, c.DbPassword))
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Fatalln("Failed to connect to db.")
	}

	api.ServeError = errors.ServeError
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UsersDeleteUsersIDHandler == nil {
		api.UsersDeleteUsersIDHandler = users.DeleteUsersIDHandlerFunc(func(params users.DeleteUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUsersID has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.UsersGetUsersIDHandler == nil {
		api.UsersGetUsersIDHandler = users.GetUsersIDHandlerFunc(func(params users.GetUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersID has not yet been implemented")
		})
	}
	if api.UsersPostUsersHandler == nil {
		api.UsersPostUsersHandler = users.PostUsersHandlerFunc(func(params users.PostUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsers has not yet been implemented")
		})
	}
	if api.UsersPutUsersIDHandler == nil {
		api.UsersPutUsersIDHandler = users.PutUsersIDHandlerFunc(func(params users.PutUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PutUsersID has not yet been implemented")
		})
	}

	handler.ConfigureApiHandlers(api, db.CreatePostgres(gorm))

	api.ServerShutdown = func() {
		gorm.Close()
		log.Info("Database connection closed successfully")
	}

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
