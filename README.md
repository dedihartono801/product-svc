
## Diagram


## Description

[Microservices with Go, gRPC and TLS server side, API Gateway, and Authentication]

## Repositories

- https://github.com/dedihartono801/product-svc - Product SVC (gRPC)
- https://github.com/dedihartono801/order-svc - Order SVC (gRPC)
- https://github.com/dedihartono801/auth-svc - Authentication SVC (gRPC)
- https://github.com/dedihartono801/api-gateway - API Gateway (HTTP)
- https://github.com/dedihartono801/protobuf - Proto file
- https://github.com/dedihartono801/ssl - For generate your ssl each service

## Installation

```bash
#create folder "ssl" in each services (auth-svc,order-svc,product-svc) 
#go to your ssl repo and generate ssl with command below:
$ ./run.sh
#and then ssl generated, copy all folders ssl (auth-svc,order-svc,product-svc) and paste to ssl folder that you have made
```

## Running the app

```bash
$ make server
```

## Generate protobuf

```bash
#go to your protobuf repo and run command below:
$ ./generator.sh
```



