package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/OSBC-LLC/togo-subgraph-main/ent"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/migrate"
	"github.com/OSBC-LLC/togo-subgraph-main/graph"
	"github.com/newrelic/go-agent/v3/newrelic"

	kit_utils "github.com/sailsforce/gomicro-kit/utils"
)

const defaultPort = "8080"

func init() {
	if err := kit_utils.InitEnv(); err != nil {
		log.Println("error loading .env: ", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var entOptions []ent.Option
	if os.Getenv("LOG_LEVEL") == "debug" || os.Getenv("LOG_LEVEL") == "silly" {
		entOptions = append(entOptions, ent.Debug())
	}

	client, err := ent.Open("postgres", kit_utils.GetDSN(os.Getenv("DATABASE_URL")), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if os.Getenv("MIGRATE") == "true" {
		if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
			log.Fatalf("failed created schema resources: %v", err)
		}
	}

	// new relic
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_APP_LICENSE")),
	)
	if err != nil {
		log.Printf("new relic init failed: %v", err)
	}

	complexity, err := strconv.Atoi(os.Getenv("QUERY_COMPLEXITY"))
	if err != nil {
		log.Printf("error getting query_complexity. Defaulting to: 15")
		complexity = 15
	}
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(extension.FixedComplexityLimit(complexity))

	if os.Getenv("LOG_LEVEL") == "silly" {
		srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetFieldContext(ctx)
			log.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			log.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
			return res, err
		})
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(newrelic.WrapHandle(newRelicApp, "/query", srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
