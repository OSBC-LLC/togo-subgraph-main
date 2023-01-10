package main

import (
	"context"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	log.Println("dagger: golangci-lint check...")
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

	linter := client.Container().From("golangci/golangci-lint:v1.50.1")
	linter = linter.WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0")

	_, err = linter.WithExec([]string{"golangci-lint", "run", "-v", "--timeout", "2m"}).Stdout(ctx)
	if err != nil {
		log.Fatalf("Error performing golangci-lint: %s\n", err)
	}

	log.Println("dagger: golangci-lint check: DONE")
}
