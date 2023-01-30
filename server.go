package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MarcoVitangeli/covid-graphql-api/dataset"
	"github.com/MarcoVitangeli/covid-graphql-api/internal/cases"
	"github.com/MarcoVitangeli/covid-graphql-api/internal/dataloader"
	"github.com/MarcoVitangeli/covid-graphql-api/internal/logger"
	"github.com/MarcoVitangeli/covid-graphql-api/internal/platform/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MarcoVitangeli/covid-graphql-api/graph"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("error loading .env file", err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.New(os.Getenv("MYSQL_DNS"))

	if err != nil {
		panic(err)
	}

	service := cases.NewService(db)

	resolver := graph.NewResolver(service)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	dService := dataloader.NewService(db)
	http.Handle("/load_dataset", dataset.HandleDataLoad(dService))

	logger.Info("starting server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
