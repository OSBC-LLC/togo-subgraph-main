build:
	@echo "Building togo-subgraph-main"
	@go build -o bin/togo-subgraph-main -v .
	@echo "done."

db:
	@echo "building docker images"
	@docker compose build
	@echo done.

du:
	@echo "starting docker images"
	@docker compose up

dp:
	@echo "starting ONLY docker postgres"
	@docker compose -f docker-compose.psql.yml up

da: db du

dagb:
	@echo "running dagger: build"
	@go run dagger/build/build.go
	@echo done.

dagl:
	@echo "running dagger: lint"
	@go run dagger/lint/lint.go
	@echo done.
