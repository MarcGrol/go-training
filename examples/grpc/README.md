# Project


## Use a grpc spec to act both grpc as rest traffic endpoint

explained here: https://github.com/grpc-ecosystem/grpc-gateway


## Setup

    go get -u github.com/golang/protobuf/protoc-gen-go
    
    cd ${GOPATH}/github.com/MarcGrol/go-training/examples/grpc
    go generate ./...
    go install ./...
    
## Start grpc server

    notificationserver # listens on port 50051

## Interact with this server using grpc
 
     notificationclient # sends large amounts of requests to localhost:50051
     
## Start rest server that proxies to grpc server

    notificationserverproxy # listens on port 8080

NB: must be started from notificationserverproxy-dir to work

##  Interact with this proxy using rest

    curl -vvv -X GET -H 'Accept: application/json' http://localhost:8080/api/notification/status/111222
    curl -vvv -X POST -H 'Content-Type: application/json' -H 'Accept: application/json' --data '{"recipientPhoneNumber":"31648928856", "body":"my body"}' http://localhost:8080/api/notification/sms
    curl -vvv -X POST -H 'Content-Type: application/json' -H 'Accept: application/json' --data '{"recipientEmailAddress":"mgrol@xebia.com", "subject": "my subject", "body":"my body"}' http://localhost:8080/api/notification/email

## The rest proxy also serves a interactive swagger-ui

Swagger ui available at: http://localhost:8080/swaggerui/

go generate creates file "spec/notification.swagger.json"

    cp spec/notification.swagger.json notificationserverproxy/swaggerui/swagger.json
    
