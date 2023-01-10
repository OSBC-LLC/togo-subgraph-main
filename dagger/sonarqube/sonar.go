package main

import (
	"context"
	"log"
	"os"

	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.dagger"); err != nil {
		log.Printf("no .env.dagger file")
	}
}

func main() {
	log.Println("dagger: sonarqube scan...")
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

	sonar := client.Container().From("sonarsource/sonar-scanner-cli")

	sonar = sonar.WithEnvVariable("SONAR_LOGIN", os.Getenv("SONAR_LOGIN"))
	sonar = sonar.WithEnvVariable("SONAR_HOST_URL", os.Getenv("SONAR_HOST_URL"))

	sonar = sonar.WithMountedDirectory("/usr/src", src).
		WithWorkdir("/usr/src")

	_, err = sonar.WithExec([]string{"sonar-scanner"}).Stdout(ctx)
	if err != nil {
		log.Fatalf("Error performing golangci-lint: %s\n", err)
	}

	log.Println("dagger: sonarqube scan: DONE")
}
