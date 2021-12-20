package main

import (
	"context"
	"time"
)

type Client struct {
	UID          string    `json:"uid"`
	FullName     string    `json:"fullName"`
	AddressLine  string    `json:"addressLine"`
	EmailAddress string    `json:"emailAddress"`
	PhoneNumber  string    `json:"phoneNumber"`
	CreatedAt    time.Time `json:"createAt"`
	LastModified time.Time `json:"lastModified"`
}

type ClientStore interface {
	Create(ctx context.Context, appointment Client) error
	Modify(ctx context.Context, appointment Client) error
	GetOnUid(ctx context.Context, appointmentUID string) (Client, bool, error)
	Search(ctx context.Context) ([]Client, error)
	Remove(ctx context.Context, userUID string) error
}
