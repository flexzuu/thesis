# benchmark

### create template for a service eg. user service
```
docker run --rm -v (PWD):/local openapitools/openapi-generator-cli generate -i /local/user.yaml -g go-server -o /local/out/go
```