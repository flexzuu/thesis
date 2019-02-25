base = ./micro-service/grpc/

user = user/user
post = post/post
rating = rating/rating

facade = facade/facade

stats = stats

generate:
	protoc -I $(base) -I $(base)$(user) $(base)$(user)/*.proto  --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(base) -I $(base)$(post) $(base)$(post)/*.proto  --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(base) -I $(base)$(rating) $(base)$(rating)/*.proto  --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(base) -I $(base)$(facade) $(base)$(facade)/*.proto  --go_out=plugins=grpc:$(GOPATH)/src
	protoc -I $(base) -I $(base)$(stats) $(base)$(stats)/*.proto  --go_out=plugins=grpc:$(GOPATH)/src
