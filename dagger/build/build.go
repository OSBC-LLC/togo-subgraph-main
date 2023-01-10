package main

import (
	"context"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	log.Println("dagger: build...")
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		log.Fatalf("Error connecting to Dagger Engine: %s\n", err)
	}

	defer client.Close()

	src := client.Host().Directory(".")
	if err != nil {
		log.Fatalf("Error getting reference to host directory: %s\n", err)
	}

	golang := client.Container().From("golang:1.19.3")
	golang = golang.WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0")

	golang = golang.WithExec([]string{"go", "build", "-o", "bin/togo-subgraph-main", "-v"})

	path := "bin/"

	build := golang.Directory(path)

	_, err = build.Export(ctx, path)
	if err != nil {
		log.Fatalf("Error writing directory: %s\n", err)
	}
	log.Println("dagger: build: DONE")
}
