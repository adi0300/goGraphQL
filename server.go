package main

import (
	"goGraphQL/config"
	"goGraphQL/db"
	"goGraphQL/graph"
	"goGraphQL/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("app.env")
	config.InitConfig()

	db.InitDatabase()

	router := chi.NewRouter()

	router.Handle("/"+config.GetConfig().AppVersion+"/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	router.Handle("/", playground.Handler("GraphQL playground", "/"+config.GetConfig().AppVersion+"/query"))
	//	http.Handle("/query", srv)

	log.Printf("connected to http://localhost:8080/")
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}
