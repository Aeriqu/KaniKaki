package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Aeriqu/kanikaki/services/api-gateway/clients"
	"github.com/Aeriqu/kanikaki/services/api-gateway/resolver"
	"github.com/Aeriqu/kanikaki/services/api-gateway/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// headerParser acts as a middleware for http.Handler and adds the headers to
// the context so it can be used later to retrieve info like authorization.
func headerParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := context.WithValue(request.Context(), "headers", request.Header)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func main() {
	schemaString := schema.GetSchemaString()
	schemaParsed := graphql.MustParseSchema(
		schemaString,
		&resolver.RootResolver{},
		graphql.UseFieldResolvers(),
	)
	log.Printf("Starting with schema:\n%q", schemaString)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: headerParser(&relay.Handler{Schema: schemaParsed}),
	}

	log.Fatal(httpServer.ListenAndServe())
	clients.CloseConnections()
}
