# Basic-API Boilerplate

This is just a boilerplate developed while learning to build and architecture a REST API server in Golang.

## Development

### Start the application in local

```bash
go mod tidy
go run app.go
```

### Up all the containers and run the program

```
make up
```

### Check more useful commands in the Makefile

```
make help
```

### Accessing the API
Now you can connect with your favourite API client in localhost:3000, modify the docker-compose file to 

## Documentation

### Architecture of the App

basic-api is a REST API app, it has been built in Go+Fiber+gORM.

### REST API

Find the OpenAPI document in .docs/openapi.yaml
