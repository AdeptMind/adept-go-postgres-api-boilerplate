package main

import (
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/config"
	"os"

	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"

	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Load config so that port is properly set
	config.LoadConfig()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewBoilerplateAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		err := server.Shutdown()
		if err != nil {
			log.WithFields(log.Fields{"err": err.Error()}).Error("Error in server shutdown")
		}
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Boilerplate"
	parser.LongDescription = "A boilerplate for a Go Api"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
