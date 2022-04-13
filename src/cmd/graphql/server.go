package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cindy1408/gym/src/cmd/graphql/graph"
	"github.com/cindy1408/gym/src/cmd/graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// _, err := NewDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// // New Database
// func NewDatabase() (*gorm.DB, error) {
// 	databaseURL := "postgresql://user:password@localhost:5432/gym?sslmode=disable"
// 	fmt.Println("INIT DATABASE")
// 	return gorm.Open(postgres.New(postgres.Config{
// 		DSN: databaseURL,
// 	}), &gorm.Config{})
// }

// // connecting to the database
// func (s *server) Connect() error {
// 	dsn := "postgresql://user:password@localhost:5432/gym?sslmode=disable"
// 	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Sprintf("cannot connect to database! error: %w", err)
// 		return nil
// 	} else {
// 		fmt.Println("Connected to database!")
// 	}
// 	return nil
// }
