package uuider

//go:generate mockgen -source=uuider.go -destination=uuiderMocks.go -package=uuider UuidGenerator

type UuidGenerator interface {
	GenerateUuid() string
}
