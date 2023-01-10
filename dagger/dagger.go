package main

import (
	"context"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		log.Fatalf("Error connecting to Dagger Engine: %s", err)
	}

	defer client.Close()

	src := client.Host().Workdir()
	if err != nil {
		log.Fatalf("Error getting reference to host directory: %s", err)
	}

	golang := client.Container().From("golang:1.19.3")
	golang = golang.WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0")

	golang = golang.Exec(
		dagger.ContainerExecOpts{
			Args: []string{"go", "build", "-o", "bin/togo-subgraph-main", "-v"},
		},
	)

	path := "bin/"

	build := golang.Directory(path)

	_, err = build.Export(ctx, path)
	if err != nil {
		log.Fatalf("Error writing directory: %s", err)
	}
}
