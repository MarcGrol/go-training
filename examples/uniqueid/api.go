package uniqueid

//go:generate mockgen -source=api.go -destination=uidGeneratorMock.go -package=uniqueid Generator
type Generator interface {
	Generate() string
}
