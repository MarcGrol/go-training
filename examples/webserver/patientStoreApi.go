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

type withinTransactionFunc func(ctx context.Context) error

type PatientStore interface {
	RunInTransaction(ctx context.Context, run withinTransactionFunc) error
	Create(ctx context.Context, appointment Patient) error
	Modify(ctx context.Context, appointment Patient) error
	GetOnUid(ctx context.Context, appointmentUID string) (Patient, bool, error)
	Search(ctx context.Context) ([]Patient, error)
	Remove(ctx context.Context, userUID string) error
}
