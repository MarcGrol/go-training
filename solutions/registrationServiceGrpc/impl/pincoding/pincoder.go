package pincoding

import (
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/pincoder"
)

type predictablePincodeGenerator struct{}

func New() pincoder.PincodeGenerator {
	return &predictablePincodeGenerator{}
}

func (pc predictablePincodeGenerator) GeneratePincode() int {
	return 1234
}
