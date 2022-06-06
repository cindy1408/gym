package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// type server struct {
// 	server *http.Server
// 	db     *gorm.DB
// 	ctx    context.Context
// }

type Server interface {
	Connect() error
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var db *gorm.DB

	db, _ = NewDatabase()

	resolver := &graph.Resolver{
		DB:     db,
	}

	// populate database tables
	if err := resolver.Init(); err != nil {
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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	// standard package to make api calls
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func NewDatabase() (*gorm.DB, error) {
	databaseURL := "host=localhost user=postgres password=password dbname=gym port=5432 sslmode=disable"
	return gorm.Open(postgres.New(postgres.Config{
		DSN: databaseURL,
	}), &gorm.Config{})
}
