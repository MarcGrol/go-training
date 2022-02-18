package uuiding

import (
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/uuider"
	"github.com/google/uuid"
)

type basicuuider struct{}

func New() uuider.UuidGenerator {
	return &basicuuider{}
}

func (u basicuuider) GenerateUuid() string {
	return uuid.NewString()
}
