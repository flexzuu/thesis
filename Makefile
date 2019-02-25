base = ./micro-service/grpc/

user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

generate:
	protoc --go_out=plugins=grpc:$(GOPATH)/src -I $(base) -I $(base)$(user) $(base)$(user)/*.proto
	protoc --go_out=plugins=grpc:$(GOPATH)/src -I $(base) -I $(base)$(post) $(base)$(post)/*.proto
	protoc --go_out=plugins=grpc:$(GOPATH)/src -I $(base) -I $(base)$(rating) $(base)$(rating)/*.proto
	protoc --go_out=plugins=grpc:$(GOPATH)/src -I $(base) -I $(base)$(facade) $(base)$(facade)/*.proto
	protoc --go_out=plugins=grpc:$(GOPATH)/src -I $(base) -I $(base)$(stats) $(base)$(stats)/*.proto
up:
	docker-compose up --build --scale client=0 --scale client-facade=0 -d
logs:
	docker-compose logs -f
down: docker-compose down
benchmark:
	docker-compose up --no-deps --build client
	docker-compose up --no-deps --build client-facade
gui:
	grpcui -port 50160 -plaintext localhost:50060 & grpcui -port 50151 -plaintext localhost:50051 & grpcui -port 50152 -plaintext localhost:50052 & grpcui -port 50153 -plaintext localhost:50053 & 