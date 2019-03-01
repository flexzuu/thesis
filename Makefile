base = micro-service/
project = github.com/flexzuu/benchmark/
grpc = grpc/
rest = rest/
grpcbase = .$(base)$(grpc)
gosrc = $(GOPATH)/src/
restbase = $(gosrc)$(project)$(base)$(rest)
user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

generate:
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(user) $(grpcbase)$(user)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(post) $(grpcbase)$(post)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(rating) $(grpcbase)$(rating)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(facade) $(grpcbase)$(facade)/*.proto
	protoc --go_out=plugins=grpc:$(gosrc) -I $(grpcbase) -I $(grpcbase)$(stats) $(grpcbase)$(stats)/*.proto
up:
	docker-compose up --build --scale client=0 --scale client-facade=0 -d
logs:
	docker-compose logs -f
down: docker-compose down

benchmark: benchmark-client benchmark-client-facade
benchmark-client:
	docker-compose up --no-deps --build client
benchmark-client-facade:
	docker-compose up --no-deps --build client-facade

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