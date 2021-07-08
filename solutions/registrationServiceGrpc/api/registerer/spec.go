package registerer

//go:generate protoc -I/usr/local/include -I . --go_out=plugins=grpc:../../impl/registration	 ./registration.proto

func init() {

}
