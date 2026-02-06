package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"gql-fiber/internal/graphql/generated"
	"gql-fiber/internal/graphql/resolver"
)

func main() {
	app := fiber.New()

	// gqlgen server
	gqlServer := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver.Resolver{},
			},
		),
	)

	// GraphQL endpoint
	app.All("/query", adaptor.HTTPHandler(gqlServer))

	// Playground
	app.Get("/", adaptor.HTTPHandler(
		playground.Handler("GraphQL Playground", "/query"),
	))

	log.Fatal(app.Listen(":8080"))
}
