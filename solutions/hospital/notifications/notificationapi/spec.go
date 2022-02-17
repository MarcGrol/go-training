package notificationapi

//go:generate protoc -I/usr/local/include -I ../.. -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./notifications.proto

//go:generate mockgen -source=notifications_grpc.pb.go -destination=notificationClientMock.go -package=notificationapi NotificationClient

func init() {

}
