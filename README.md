# Tracker App

## Introduction

This is a containerized Golang app that can track and save incoming requests.

## Prerequisites

- Golang 1.20 version and greater
- PostgreSQL
- Docker
- Docker Compose
- Open Api v3

- This app includes openapi.yaml specification that you can use with [Swagger Editor](https://editor.swagger.io) copy and pase the spec to this web site to see all available endpoints and methods.

## Getting Started

NOTE! I strongly recommend you to run this app on MacBooks with M1, M2 etc. processor. Command for building and running application on other platform are not tested but they are defined in the project. If you will be able to run under other platform with some fixes I'm ready to REVIEW PULL REQUESTS.

For more details see corresponding titles for each command and step

### Generating models from openapi.yaml

1. First, run the following command to generate any necessary files:

```bash
go install -a
```

then

```bash
make generate
```

### Building the app

2. Then, we need to build 2 docker images to run it with docker-compose

- for main tracker application which serves HTTP requests
- to make a request to service you can use openapi.yaml specification
- first create an account with specified balance
- then create a transactions with specified account id from 1st stage

```bash
make buildx
```

- for post-processing worker which according to task should
- every N minutes 10 latest odd records must be canceled and balance should be corrected by the application.
- by default N is set to 1 minute
- to set minutes go to cmd/cancel-transaction-worker/main.go where s.Every(<minutes>).Minute() <minutes> to N

```bash
make buildx
```

### Running the App

3. To launch, we need to run docker-compose, this will launch PostgreSQL database and apply migrations to start sending requests

```bash
make composeupx
```

### Stopping the App

4. To stop, we need to run docker-compose, this will stop application and database

```bash
make composedownx
```

## Building the App

To build the app, run:

- For MacOS under M1, M2 processors (*tested on MacBook Air M2*)

```bash
make buildx
```

This will build a Docker image tagged tracker:latest based on the Dockerfile at docker/macos/arm64/Dockerfile.

- For other platforms (not tested due to lack of other platforms on hand)

```bash
make build
```

This will build a Docker image tagged tracker:latest based on the Dockerfile at docker/other/Dockerfile.

## Running the App without Database (see bellow to run with docker-compose that includes all dependencies)

To run the app, run:

Similar to previews commands
- For other platforms

```bash
make run
```

- For MacOS under M1, M2 processors (*tested on MacBook Air M2*)

```bash
make runx
```

This will start a Docker container running the tracker:latest image, which listens on port 8080.

## Using Docker Compose

If you have Docker Compose installed, you can use it to start the app and its dependencies using the following commands:

Similar to previews commands
- For other platforms

```bash
make composeup
```

- For MacOS under M1, M2 processors (*tested on MacBook Air M2*)

```bash
make composeupx
```

This will start the app and a PostgreSQL database container.

To stop the app and the database and application, run:

Similar to previews commands
- For other platforms

```bash
make composedown
```

- For MacOS under M1, M2 processors (*tested on MacBook Air M2*)

```bash
make composedownx
```

## Conclusion

That's it! You now have a containerized Golang app that can track and save incoming requests.