package appointmentapi

//go:generate protoc -I/usr/local/include -I ../.. -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./appointments.proto

//go:generate mockgen -source=appointments_grpc.pb.go -destination=appointmentClientMock.go -package=appointmentapi AppointmentInternalClient,AppointmentExternalClient

func init() {

}
