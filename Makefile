base = example/
project = github.com/flexzuu/thesis/
grpc = grpc/
openapi = openapi/
hal = hal/
graphql = graphql/
mixed = mixed/
grpcbase = .$(base)$(grpc)
gosrc = $(GOPATH)/src/
grpcbase = $(gosrc)$(project)$(base)$(grpc)
restbase = $(gosrc)$(project)$(base)$(openapi)
halbase = $(gosrc)$(project)$(base)$(hal)
graphqlbase = $(gosrc)$(project)$(base)$(graphql)
mixedbase = $(gosrc)$(project)$(base)$(mixed)
user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

docc-grpc = $(grpcbase)docker-compose.yaml
docc-openapi = $(restbase)docker-compose.yaml
docc-hal = $(halbase)docker-compose.yaml
docc-graphql = $(graphqlbase)docker-compose.yaml
docc-mixed = $(mixedbase)docker-compose.yaml

generate:
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(user) $(grpcbase)$(user)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(post) $(grpcbase)$(post)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(rating) $(grpcbase)$(rating)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(facade) $(grpcbase)$(facade)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(stats) $(grpcbase)$(stats)/*.proto
up-grpc:
	docker-compose -f $(docc-grpc) up --build --scale client=0 --scale client-facade=0 -d
logs-grpc:
	docker-compose -f $(docc-grpc) logs -f
down-grpc: 
	docker-compose -f $(docc-grpc) down
thesis-grpc: thesis-client-grpc thesis-client-facade-grpc
thesis-client-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client
thesis-client-facade-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client-facade

up-openapi:
	docker-compose -f $(docc-openapi) up --build --scale client=0 --scale client-facade=0 -d
logs-openapi:
	docker-compose -f $(docc-openapi) logs -f
down-openapi: 
	docker-compose -f $(docc-openapi) down
thesis-openapi: thesis-client-openapi thesis-client-facade-openapi
thesis-client-openapi:
	docker-compose -f $(docc-openapi) up --no-deps --build client
thesis-client-facade-openapi:
	docker-compose -f $(docc-openapi) up --no-deps --build client-facade

up-hal:
	docker-compose -f $(docc-hal) up --build --scale client=0 --scale client-facade=0 -d
logs-hal:
	docker-compose -f $(docc-hal) logs -f
down-hal: 
	docker-compose -f $(docc-hal) down
thesis-hal: thesis-client-hal thesis-client-facade-hal
thesis-client-hal:
	docker-compose -f $(docc-hal) up --no-deps --build client
# thesis-client-facade-hal:
# 	docker-compose -f $(docc-hal) up --no-deps --build client-facade

up-graphql:
	docker-compose -f $(docc-graphql) up --build --scale client=0 --scale client-facade=0 -d
logs-graphql:
	docker-compose -f $(docc-graphql) logs -f
down-graphql: 
	docker-compose -f $(docc-graphql) down
thesis-graphql: thesis-client-graphql thesis-client-facade-graphql
thesis-client-graphql:
	docker-compose -f $(docc-graphql) up --no-deps --build client
thesis-client-facade-graphql:
	docker-compose -f $(docc-graphql) up --no-deps --build client-facade

up-mixed:
	docker-compose -f $(docc-mixed) up --build --scale client=0 --scale client-facade=0 -d
logs-mixed:
	docker-compose -f $(docc-mixed) logs -f
down-mixed: 
	docker-compose -f $(docc-mixed) down
thesis-mixed: thesis-client-mixed thesis-client-facade-mixed
thesis-client-mixed:
	docker-compose -f $(docc-mixed) up --no-deps --build client
thesis-client-facade-mixed:
	docker-compose -f $(docc-mixed) up --no-deps --build client-facade

gui: gui-1 gui-2 gui-3 gui-4
gui-1: 
	grpcui -port 50160 -plaintext localhost:50060
gui-2:
	grpcui -port 50151 -plaintext localhost:50051
gui-3:
	grpcui -port 50152 -plaintext localhost:50052
gui-4:
	grpcui -port 50153 -plaintext localhost:50053 
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