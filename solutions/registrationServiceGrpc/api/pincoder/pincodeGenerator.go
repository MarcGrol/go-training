package pincoder

//go:generate mockgen -source=pincodeGenerator.go -destination=pincodeGeneratorMocks.go -package=pincoder PincodeGenerator

type PincodeGenerator interface {
	GeneratePincode() int
}
