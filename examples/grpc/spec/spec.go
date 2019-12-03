package notification

//go:generate protoc -I/usr/local/include -I ../spec --go_out=plugins=grpc:../spec --grpc-gateway_out=logtostderr=true:../spec --swagger_out=logtostderr=true:../specg ../spec/notification.proto

func init() {

}
