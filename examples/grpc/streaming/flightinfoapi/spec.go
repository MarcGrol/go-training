package flightinfoapi

//go:generate protoc -I/usr/local/include -I . --go_out=plugins=grpc:.  ./flightinfo.proto

func init() {

}
