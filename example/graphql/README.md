# GraphQL Example

## Services:
- facade
- rating
- user
- post

## Clients:
- seed: creates example data
- client: fetches data from services directly
- client-facade: fetches data from the facade

## Specialties:
- facade: exposes a general purpose data graph to solve the same need the other facade solve

## Used Commands & Tools:
- `https://github.com/99designs/gqlgen`
- `gqlgen init`
- `gqlgen`

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
├── client
│   ├── client.go
│   ├── postclient
│   │   └── postclient.go
│   ├── ratingclient
│   │   └── ratingclient.go
│   └── userclient
│       └── userclient.go
├── client-facade
│   └── client.go
├── docker-compose.yaml
├── facade
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── postclient
│   │   └── postclient.go
│   ├── ratingclient
│   │   └── ratingclient.go
│   ├── resolver.go
│   ├── schema.graphql
│   ├── server
│   │   └── server.go
│   └── userclient
│       └── userclient.go
├── post
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── models_gen.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   ├── resolver.go
│   ├── schema.graphql
│   ├── server
│   │   └── server.go
│   └── userclient
│       └── userclient.go
├── rating
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── models_gen.go
│   ├── postclient
│   │   └── postclient.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   ├── resolver.go
│   ├── schema.graphql
│   └── server
│       └── server.go
├── seed
│   ├── main.go
│   ├── postclient
│   │   └── postclient.go
│   ├── ratingclient
│   │   └── ratingclient.go
│   └── userclient
│       └── userclient.go
├── user
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── models_gen.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   ├── resolver.go
│   ├── schema.graphql
│   └── server
│       └── server.go
└── util
    └── id.go
```