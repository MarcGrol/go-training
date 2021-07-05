package main

import (
	"context"
	"time"
)

type Patient struct {
	UID          string    `json:"uid"`
	FullName     string    `json:"fullName"`
	AddressLine  string    `json:"addressLine"`
	Allergies    []string  `json:"allergies"`
	CreatedAt    time.Time `json:"createAt"`
	LastModified time.Time `json:"lastModified"`
}

type PatientStore interface {
	Put(ctx context.Context, appointment Patient) (Patient, error)
	GetOnUid(ctx context.Context, appointmentUID string) (Patient, bool, error)
	Search(ctx context.Context) ([]Patient, error)
	Remove(ctx context.Context, userUID string) error
}
