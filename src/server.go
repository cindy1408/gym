package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/cindy1408/gym/src/graphql/graph/postgres"
	"github.com/cindy1408/gym/src/graphql/graph/resolver"
	"gorm.io/gorm"
)

const defaultPort = "8080"

type Server interface {
	Connect() error
}

var DB *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	DB, _ = postgres.NewDatabase()

	pgRepo := postgres.NewRepo(DB)

	resolver := &resolver.Resolver{}

	// populate database tables
	if err := pgRepo.Init(); err != nil {
		log.Fatal("failed to create database: %v", err)
	}

	q := resolver.Query()
	m := resolver.Mutation()

	var ctx context.Context

	_, err := q.HydrateBaseExercise(ctx)
	if err != nil {
		fmt.Println("Hydrate base exercise failed")
	}

	_, err = m.HydrateMuscleGroups(ctx)
	if err != nil {
		fmt.Println("Hydrate group exercise failed")
	}

	// Remember to pass the initialised database to the server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	// standard package to make api calls
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
