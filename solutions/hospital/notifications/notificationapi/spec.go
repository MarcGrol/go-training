package notificationapi

//go:generate protoc -I/usr/local/include -I  ../.. -I . --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. ./notifications.proto

//go:generate mockgen -source=notifications.pb.go -destination=notificationClientMock.go -package=notificationapi NotificationClient

func init() {

}
