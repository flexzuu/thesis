# Function / gRPC Example

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

- stats: implemented by multiple services to allow to count the number of round-trips

## Used Commands & Tools:

- `protoc`
 
 ```bash
protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(user) $(grpcbase)$(user)/*.proto
protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(post) $(grpcbase)$(post)/*.proto
protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(rating) $(grpcbase)$(rating)/*.proto
protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(facade) $(grpcbase)$(facade)/*.proto
protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(stats) $(grpcbase)$(stats)/*.proto
```

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
│   └── client.go
├── client-facade
│   └── client.go
├── docker-compose.yaml
├── facade
│   ├── facade
│   │   ├── facade.pb.go
│   │   └── facade.proto
│   └── server
│       └── main.go
├── post
│   ├── post
│   │   ├── post.pb.go
│   │   └── post.proto
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   └── server
│       ├── conversion.go
│       └── main.go
├── rating
│   ├── rating
│   │   ├── rating.pb.go
│   │   └── rating.proto
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   └── server
│       ├── convert.go
│       └── main.go
├── seed
│   └── main.go
├── stats
│   ├── stats.go
│   ├── stats.pb.go
│   └── stats.proto
└── user
    ├── repo
    │   ├── entity
    │   ├── inmemmory
    │   └── repo.go
    ├── server
    │   ├── convert.go
    │   └── main.go
    └── user
        ├── user.pb.go
        └── user.proto
```