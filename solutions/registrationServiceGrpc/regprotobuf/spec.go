package regprotobuf

//go:generate protoc -I/usr/local/include -I ../.. -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./registration.proto

const (
	DefaultPort = ":60002"
)

func init() {}
