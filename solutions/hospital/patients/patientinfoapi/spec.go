package patientinfoapi

//go:generate protoc -I/usr/local/include -I  ../.. -I . --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. ./patients.proto

//go:generate mockgen -source=patients.pb.go -destination=patientinfoClientMock.go -package=patientinfoapi PatientInfoClient

func init() {

}
