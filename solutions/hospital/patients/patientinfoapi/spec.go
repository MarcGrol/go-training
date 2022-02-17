package patientinfoapi

//go:generate protoc -I/usr/local/include -I  ../.. -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./patients.proto

//go:generate mockgen -source=patients_grpc.pb.go -destination=patientinfoClientMock.go -package=patientinfoapi PatientInfoClient

func init() {

}
