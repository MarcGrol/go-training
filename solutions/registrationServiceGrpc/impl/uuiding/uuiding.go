package uuiding

import (
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/uuider"
	"github.com/google/uuid"
)

type basicuuider struct{}

func New() uuider.UuidGenerator {
	return &basicuuider{}
}

func (u basicuuider) GenerateUuid() string {
	return uuid.NewString()
}
