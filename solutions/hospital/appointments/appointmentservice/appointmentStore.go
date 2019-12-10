package main

import (
	"sync"
)

type AppointmentStatus int

const (
	AppointmentStatusUnknown AppointmentStatus = 0
	AppointmentStatusRequested
	AppointmentStatusConfirmed
)

type Appointment struct {
	AppointmentUID string
	UserUID        string
	DateTime       string
	Location       string
	Topic          string
	Status         AppointmentStatus
}

type AppointmentStore interface {
	PutAppointment(appointment Appointment) (Appointment, error)
	GetAppointmentOnUid(appointmentUID string) (Appointment, bool, error)
	GetAppointmentsOnUserUid(userUID string) ([]Appointment, error)
	GetAppointmentsOnStatus(status AppointmentStatus) ([]Appointment, error)
}

type appointmentStore struct {
	sync.Mutex
	uider        Uider
	appointments map[string]Appointment
}

func newAppointmentStore(uider Uider) AppointmentStore {
	return &appointmentStore{
		uider: uider,
		appointments: map[string]Appointment{
			"a": {AppointmentUID: "a", UserUID: "1", Location: "Leuven", Topic: "onderzoek", Status: AppointmentStatusRequested},
			"b": {AppointmentUID: "b", UserUID: "2", Location: "Leuven", Topic: "scan", Status: AppointmentStatusConfirmed},
		},
	}
}

func (as *appointmentStore) PutAppointment(appointment Appointment) (Appointment, error) {
	as.Lock()
	defer as.Unlock()

	if appointment.AppointmentUID == "" {
		appointment.AppointmentUID = as.uider.Create()
	}
	as.appointments[appointment.AppointmentUID] = appointment
	return appointment, nil
}

func (as *appointmentStore) GetAppointmentOnUid(appointmentUID string) (Appointment, bool, error) {
	as.Lock()
	defer as.Unlock()

	patient, found := as.appointments[appointmentUID]
	return patient, found, nil
}

func (as *appointmentStore) GetAppointmentsOnUserUid(userUID string) ([]Appointment, error) {
	as.Lock()
	defer as.Unlock()

	found := []Appointment{}
	for _, appointment := range as.appointments {
		if appointment.UserUID == userUID {
			found = append(found, appointment)
		}
	}
	return found, nil
}

func (as *appointmentStore) GetAppointmentsOnStatus(status AppointmentStatus) ([]Appointment, error) {
	as.Lock()
	defer as.Unlock()

	found := []Appointment{}
	for _, appointment := range as.appointments {
		if appointment.Status == status {
			found = append(found, appointment)
		}
	}
	return found, nil
}
