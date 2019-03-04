base = micro-service/
project = github.com/flexzuu/benchmark/
grpc = grpc/
rest = rest/
grpcbase = .$(base)$(grpc)
gosrc = $(GOPATH)/src/
grpcbase = $(gosrc)$(project)$(base)$(grpc)
restbase = $(gosrc)$(project)$(base)$(rest)
user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

docc-grpc = $(grpcbase)docker-compose.yaml
docc-rest = $(restbase)docker-compose.yaml

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
down-grpc: docker-compose -f $(docc-grpc) down
benchmark-grpc: benchmark-client-grpc benchmark-client-facade-grpc
benchmark-client-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client
benchmark-client-facade-grpc:
	docker-compose -f $(docc-grpc) up --no-deps --build client-facade

up-rest:
	docker-compose -f $(docc-rest) up --build --scale client=0 --scale client-facade=0 -d
logs-rest:
	docker-compose -f $(docc-rest) logs -f
down-rest: docker-compose -f $(docc-rest) down
benchmark-rest: benchmark-client-rest benchmark-client-facade-rest
benchmark-client-rest:
	docker-compose -f $(docc-rest) up --no-deps --build client
benchmark-client-facade-rest:
	docker-compose -f $(docc-rest) up --no-deps --build client-facade


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