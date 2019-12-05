package appointmentapi

//go:generate protoc -I/usr/local/include -I ../.. -I . --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. ./appointments.proto

func init() {

}
