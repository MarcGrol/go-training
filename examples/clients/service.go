package main

import (
	"context"
	"fmt"
)

type clientWebService struct {
	uuider      func() string
	clientStore ClientStore
	emailSender EmailSender
}

func NewPatientService(uuider func() string, clientStore ClientStore, emailSender EmailSender) *clientWebService {
	return &clientWebService{
		uuider:      uuider,
		clientStore: clientStore,
		emailSender: emailSender,
	}
}

func (s *clientWebService) getPatientOnUID(c context.Context, clientUID string) (Client, bool, error) {
	return s.clientStore.GetOnUid(c, clientUID)
}

func (s *clientWebService) createClient(c context.Context, client Client) (Client, error) {
	client.UID = s.uuider()

	// TODO Send welcome email when email-address is not empty

	// TODO Send welcome sms when phone-Number is not empty

	err := s.clientStore.Create(c, client)
	return client, err
}

func (s *clientWebService) modifyClientOnUid(c context.Context, client Client) error {
	_, found, err := s.getPatientOnUID(c, client.UID)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Not found")
	}

	return s.clientStore.Modify(c, client)
}

func (s *clientWebService) deleteClientOnUid(c context.Context, clientUID string) error {
	_, found, err := s.getPatientOnUID(c, clientUID)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Not found")
	}
	return s.clientStore.Remove(c, clientUID)
}
