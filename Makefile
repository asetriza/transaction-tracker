generate:
	go generate ./...

buildx:
	docker build --platform=linux/amd64 -t tracker:latest . -f docker/macos/arm64/Dockerfile

runx:
	docker run -d -p 8080:8080 --platform=linux/amd64 tracker:latest

build:
	docker build -t tracker:latest . -f docker/other/Dockerfile

run:
	docker run -d -p 8080:8080 --platform=linux/amd64 tracker:latest

composeupx:
	docker-compose --file docker-compose/macos/arm64/docker-compose.yaml up

composedownx:
	docker-compose --file docker-compose/macos/arm64/docker-compose.yaml down

composeup:
	docker-compose --file docker-compose/other/docker-compose.yaml up

composedown:
	docker-compose --file docker-compose/other/docker-compose.yaml down