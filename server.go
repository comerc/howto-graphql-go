package main

import (
	"example/internal/gql/resolver"
	"example/internal/gql/runtime"
	"example/internal/pkg/db/mysql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8888"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mysql.InitDB()
	defer mysql.CloseDB()
	mysql.Migrate()

	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(runtime.Config{Resolvers: resolver.NewRootResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
