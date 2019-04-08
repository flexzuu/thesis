# RESTish / OpenAPi Example

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

- `openapitools/openapi-generator-cli`
 
 ```bash
generate-openapi-user:
	docker run --rm -v $(restbase)user/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/user.yaml -g go-gin-server -o /local/out/go
generate-openapi-user-client:
	# client
	docker run --rm -v $(restbase)user/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/user.yaml -g go -DpackageName=client -o /local/client
generate-openapi-post:
	# server
	docker run --rm -v $(restbase)post/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/post.yaml -g go-gin-server -o /local/out/go
generate-openapi-post-client:
	# client
	docker run --rm -v $(restbase)post/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/post.yaml -g go -DpackageName=client -o /local/client
generate-openapi-rating:
	# server
	docker run --rm -v $(restbase)rating/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/rating.yaml -g go-gin-server -o /local/out/go
generate-openapi-rating-client:
	# client
	docker run --rm -v $(restbase)rating/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/rating.yaml -g go -DpackageName=client -o /local/client
generate-openapi-facade:
	# server
	docker run --rm -v $(restbase)facade/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/facade.yaml -g go-gin-server -o /local/out/go
generate-openapi-facade-client:
	# client
	docker run --rm -v $(restbase)facade/openapi:/local openapitools/openapi-generator-cli generate --model-name-suffix=Model -i /local/facade.yaml -g go -DpackageName=client -o /local/client
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
│   ├── openapi
│   │   ├── api_facade.go
│   │   ├── client
│   │   ├── facade.yaml
│   │   ├── model_author_detail_model.go
│   │   ├── model_post_detail_model.go
│   │   ├── model_post_list_model.go
│   │   └── routers.go
│   └── server
│       └── main.go
├── post
│   ├── openapi
│   │   ├── api_post.go
│   │   ├── client
│   │   ├── model_create_post_model.go
│   │   ├── model_post_list_model.go
│   │   ├── model_post_model.go
│   │   ├── post.yaml
│   │   └── routers.go
│   ├── repo
│   │   ├── entity
│   │   ├── inmemmory
│   │   └── repo.go
│   └── server
│       └── main.go
├── rating
│   ├── openapi
│   │   ├── api_rating.go
│   │   ├── client
│   │   ├── model_create_rating_model.go
│   │   ├── model_rating_list_model.go
│   │   ├── model_rating_model.go
│   │   ├── rating.yaml
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
    ├── openapi
    │   ├── api_user.go
    │   ├── client
    │   ├── model_create_user_model.go
    │   ├── model_user_model.go
    │   ├── routers.go
    │   └── user.yaml
    ├── repo
    │   ├── entity
    │   ├── inmemmory
    │   └── repo.go
    └── server
        └── main.go
```