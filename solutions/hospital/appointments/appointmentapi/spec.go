package appointmentapi

//go:generate protoc -I/usr/local/include -I ../.. -I . --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. ./appointments.proto

//go:generate mockgen -source=appointments.pb.go -destination=appointmentClientMock.go -package=appointmentapi AppointmentInternalClient,AppointmentExternalClient

func init() {

}
