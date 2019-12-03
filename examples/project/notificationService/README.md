# Project

## Setup

    go get -u github.com/golang/protobuf/protoc-gen-go
    
    go generate ./...
    

## Submit and fetch using grpc
 
     notificationclient
    
## Fetch using rest

     curl -vvv -X GET -H 'Accept: application/json' http://localhost:8080/api/notification/status/111222

## Submit sms using rest

    curl -vvv -X POST -H 'Content-Type: application/json' -H 'Accept: application/json' --data '{"recipientPhoneNumber":"31648928856", "body":"my body"}' http://localhost:8080/api/notification/sms

## Submit email using rest

    curl -vvv -X POST -H 'Content-Type: application/json' -H 'Accept: application/json' --data '{"recipientEmailAddress":"mgrol@xebia.com", "subject": "my subject", "body":"my body"}' http://localhost:8080/api/notification/email
    