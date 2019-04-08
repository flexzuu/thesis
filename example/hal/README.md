# RESTful / HAL Example

## Services:
- facade
- rating
- user
- post

## Clients:
- seed: creates example data
- client: fetches data from services directly
- client-facade: fetches data from the facade

## Used Commands & Tools:

- `https://github.com/leibowitz/halgo`
- openapi example as starting point

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
│   ├── api
│   │   ├── api_facade.go
│   │   ├── model_author_detail_model.go
│   │   ├── model_post_detail_model.go
│   │   ├── model_post_list_model.go
│   │   └── routers.go
│   └── server
│       └── main.go
├── post
│   ├── api
│   │   ├── api_post.go
│   │   ├── model_create_post_model.go
│   │   ├── model_post_list_model.go
│   │   ├── model_post_model.go
│   │   └── routers.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   └── server
│       └── main.go
├── rating
│   ├── api
│   │   ├── api_rating.go
│   │   ├── model_create_rating_model.go
│   │   ├── model_rating_list_model.go
│   │   ├── model_rating_model.go
│   │   └── routers.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   └── server
│       └── main.go
├── seed
│   └── main.go
└── user
    ├── api
    │   ├── api_user.go
    │   ├── model_create_user_model.go
    │   ├── model_user_model.go
    │   └── routers.go
    ├── repo
    │   ├── entity
    │   ├── inmemmory
    │   └── repo.go
    └── server
        └── main.go
```