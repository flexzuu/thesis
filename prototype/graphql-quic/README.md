# GraphQL over HTTP/3 - QUIC

### Test HTTP 
```bash
go run ./server/http &
go run ./server/quic &
go run ./client/ -method http |jq ".data.method" # returns http
go run ./client/ -method quic |jq ".data.method" # returns quic
```


```
├── README.md
├── client
│   └── client.go
├── graphql
│   └── main.go
├── model
│   └── model.go
├── server
│   ├── http
│   └── quic
└── testdata
    ├── ca.pem
    ├── cert.go
    ├── cert.pem
    ├── cert_test.go
    ├── priv.key
    └── testdata_suite_test.go
```
    
