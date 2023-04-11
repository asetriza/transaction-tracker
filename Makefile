install:
	go install -a

generate:
	go generate ./...

test:
	go test ./...

buildx:
	docker build --platform=linux/amd64 -t tracker:latest . -f docker/macos/arm64/Dockerfile

buildxworker:
	docker build --platform=linux/amd64 -t cancel-transaction-worker:latest . -f docker/macos/arm64/cancel-transaction-worker/Dockerfile

runx:
	docker run -d -p 8080:8080 --platform=linux/amd64 tracker:latest

build:
	docker build -t tracker:latest . -f docker/other/Dockerfile

buildworker:
	docker build -t cancel-transaction-worker:latest . -f docker/other/cancel-transaction-worker/Dockerfile

run:
	docker run -d -p 8080:8080 --platform=linux/amd64 tracker:latest

buildprojectx: install generate test buildx buildxworker

buildproject: install generate test build buildworker

composeupx:
	docker-compose --file docker-compose/macos/arm64/docker-compose.yaml up

composedownx:
	docker-compose --file docker-compose/macos/arm64/docker-compose.yaml down

composeup:
	docker-compose --file docker-compose/other/docker-compose.yaml up

composedown:
	docker-compose --file docker-compose/other/docker-compose.yaml down