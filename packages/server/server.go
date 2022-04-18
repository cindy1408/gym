package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cindy1408/gym/packages/config"
	"github.com/cindy1408/gym/src/graphql/graph"
	"github.com/cindy1408/gym/src/graphql/graph/generated"
	"github.com/go-chi/chi"
	"go.opencensus.io/plugin/ochttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type server struct {
	conf *config.Config
	serv *http.Server
	db   *gorm.DB
}
type Server interface {
	Connect() error
	ListenAndServe() error
	Shutdown(context.Context) error
}

func NewServer() Server {

	conf := config.MustLoadConfig()
	r := chi.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	// standard package to make api calls
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// http.HandleFunc("/aboutMe", aboutMe)

	// Docker for postgres

	var db *gorm.DB
	var err error
	db, err = NewDatabase(conf)
	if err != nil {
		log.Fatal(err)
	}

	resolver := &graph.Resolver{
		Config: conf,
		DB:     db,
	}
	// populate database tables
	if err := resolver.Init(); err != nil {
		log.Fatal("failed to create database: %v", err)
	}

	return &server{
		conf: conf,
		serv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Handler: r,
		},
		db: db,
	}

}

// New Database
func NewDatabase(conf *config.Config) (*gorm.DB, error) {
	databaseURL := "postgresql://user:password@localhost:5432/gym?sslmode=disable"
	fmt.Println("INIT DATABASE")
	return gorm.Open(postgres.New(postgres.Config{
		DSN: databaseURL,
	}), &gorm.Config{})
}

// connecting to the database
func (s *server) Connect() error {
	dsn := "postgresql://user:password@localhost:5432/gym?sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("cannot connect to database! error: %w", err)
		return nil
	} else {
		fmt.Println("Connected to database!")
	}
	return nil
}

func (s *server) ListenAndServe() error {
	log.Printf("listening on %s", s.serv.Addr)
	if s.conf.GraphQLPlayground {
		log.Printf("connect to http://%s/graphql or http://localhost:%s/graphql for GraphQL playground", s.serv.Addr, s.conf.Port)
	}
	return http.ListenAndServe(s.serv.Addr, &ochttp.Handler{
		Handler: s.serv.Handler,
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	db, err := s.db.DB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	return s.serv.Shutdown(ctx)
}
