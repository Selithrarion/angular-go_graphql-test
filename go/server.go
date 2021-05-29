package main

import (
	"github.com/selithrarion/angular-go_graphql-test/internal/auth"
	"github.com/selithrarion/angular-go_graphql-test/internal/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/selithrarion/angular-go_graphql-test/graph"
	"github.com/selithrarion/angular-go_graphql-test/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
