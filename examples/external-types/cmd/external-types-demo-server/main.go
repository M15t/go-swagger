// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/M15t/go-swagger/examples/external-types/restapi"
	"github.com/M15t/go-swagger/examples/external-types/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExternalTypesDemoAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "external types imports: external anonymous types"
	parser.LongDescription = "This sample specification exercises external types, with both x-go-type in definitions and inlined.\n\nIt demonstrates how to use the x-go-type extension to plug external type definitions in the generated code,\nfor models (e.g. for properties, arrays or maps) or operations.\n\nNotice that x-go-type works for schemas and is not supported for simple swagger types,\nused for response headers and query & path parameters.\n"
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
