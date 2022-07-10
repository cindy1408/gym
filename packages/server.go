package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cindy1408/gym/packages/api/graphql/graph/generated"
	"github.com/cindy1408/gym/packages/api/graphql/graph/resolver"
	"github.com/cindy1408/gym/packages/api/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

type Server interface {
	Connect() error
}

var db *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, _ = postgres.NewDatabase()

	pgRepo := postgres.NewRepo(db)

	resolver := &resolver.Resolver{
		PgRepo: pgRepo,
	}

	// populate database tables
	if err := pgRepo.Init(db); err != nil {
		log.Fatal("failed to create database: %v", err)
	}

	resolver.Query()
	resolver.Mutation()

	// Remember to pass the initialised database to the server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	// standard package to make api calls
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
