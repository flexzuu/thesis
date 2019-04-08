# Mixed Example: gRPC for service plus GraphQL facade 

## Services:
- facade (GraphQL)
- rating (gRPC reused from ../grpc folder)
- user (gRPC reused from ../grpc folder)
- post (gRPC reused from ../grpc folder)

## Clients:
- seed: creates example data (reused from ../grpc folder)
- client-facade: fetches data from the facade

## Used Commands & Tools:

- reused gRPC example and the tools used for the GraphQL example

## Start instructions

- switch in directory
- start services: (`docker-compose up --build --scale client=0 --scale client-facade=0 -d`)
- show logs in new console window: (`docker-compose logs -f`)
- give it some time to boot up and execute the seeds
- run clients
  - normal: (`docker-compose up --no-deps --build client`)
  - facade: (`docker-compose up --no-deps --build client-facade`)
- make sure to shut it down again because the different example use overlapping ports
  - shutdown (`docker-compose down -v`)
  - check nothing is running (`docker ps` or `docker-compose ps`)
  
```
├── README.md
├── client-facade
│   └── client.go
├── docker-compose.yaml
├── facade
│   ├── convert.go
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── resolver.go
│   ├── schema.graphql
│   └── server
│       └── server.go
└── util
    └── id.go
```