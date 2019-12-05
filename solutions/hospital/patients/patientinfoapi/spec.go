package patientinfoapi

//go:generate protoc -I/usr/local/include -I  ../.. -I . --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. ./patients.proto

func init() {

}
