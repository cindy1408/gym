package main

import (
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
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

type server struct {
	server http.Server
	db     gorm.DB
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	// standard package to make api calls
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	_, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

}

func Connect() (gorm.DB, error) {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"user",
		"password",
		"host",
		"port",
		"gym")

	// connect to the postgres db just to be able to run the create db statement
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return *db, err
	}
	fmt.Println("successful!")
	return *db, nil
}
