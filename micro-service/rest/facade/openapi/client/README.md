# Go API client for client

a facade service

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorApi* | [**AuthorDetail**](docs/AuthorApi.md#authordetail) | **Get** /facade/author/{id} | Author Detail
*FacadeApi* | [**AuthorDetail**](docs/FacadeApi.md#authordetail) | **Get** /facade/author/{id} | Author Detail
*FacadeApi* | [**ListPosts**](docs/FacadeApi.md#listposts) | **Get** /facade/post | List Posts
*FacadeApi* | [**PostDetail**](docs/FacadeApi.md#postdetail) | **Get** /facade/post/{id} | Post Detail
*PostApi* | [**ListPosts**](docs/PostApi.md#listposts) | **Get** /facade/post | List Posts
*PostApi* | [**PostDetail**](docs/PostApi.md#postdetail) | **Get** /facade/post/{id} | Post Detail
*UserApi* | [**AuthorDetail**](docs/UserApi.md#authordetail) | **Get** /facade/author/{id} | Author Detail


## Documentation For Models

 - [AuthorDetailModel](docs/AuthorDetailModel.md)
 - [PostDetailModel](docs/PostDetailModel.md)
 - [PostListModel](docs/PostListModel.md)
 - [PostModel](docs/PostModel.md)
 - [UserModel](docs/UserModel.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author


