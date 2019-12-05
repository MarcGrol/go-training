package main

import "github.com/google/uuid"

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
	appointments map[string]Appointment
}

func newAppointmentStore() AppointmentStore {
	return &appointmentStore{
		appointments: map[string]Appointment{
			"a": {AppointmentUID: "a", UserUID: "1", Location: "Leuven", Topic: "onderzoek", Status: AppointmentStatusRequested},
			"b": {AppointmentUID: "b", UserUID: "2", Location: "Leuven", Topic: "scan", Status: AppointmentStatusConfirmed},
		},
	}
}

func (as *appointmentStore) PutAppointment(appointment Appointment) (Appointment, error) {
	if appointment.AppointmentUID == "" {
		appointment.AppointmentUID = uuid.New().String() // TODO Mock
	}
	as.appointments[appointment.AppointmentUID] = appointment
	return appointment, nil
}

func (as appointmentStore) GetAppointmentOnUid(appointmentUID string) (Appointment, bool, error) {
	patient, found := as.appointments[appointmentUID]
	return patient, found, nil
}

func (as appointmentStore) GetAppointmentsOnUserUid(userUID string) ([]Appointment, error) {
	found := []Appointment{}
	for _, appointment := range as.appointments {
		if appointment.UserUID == userUID {
			found = append(found, appointment)
		}
	}
	return found, nil
}

func (as appointmentStore) GetAppointmentsOnStatus(status AppointmentStatus) ([]Appointment, error) {
	found := []Appointment{}
	for _, appointment := range as.appointments {
		if appointment.Status == status {
			found = append(found, appointment)
		}
	}
	return found, nil
}
