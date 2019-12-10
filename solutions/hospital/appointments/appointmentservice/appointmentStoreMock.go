package main

import (
	"errors"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentservice/appointmentstore"
)

type mockAppointmentStore struct {
	err          error
	putError     bool
	exists       bool
	appointment  appointmentstore.Appointment
	appointments []appointmentstore.Appointment
}

func NewErrorMockAppointmentStore(err error) appointmentstore.AppointmentStore {
	return &mockAppointmentStore{
		err:          err,
		putError:     false,
		appointment:  appointmentstore.Appointment{},
		appointments: []appointmentstore.Appointment{},
	}
}

func NewNotFoundMockAppointmentStore() appointmentstore.AppointmentStore {
	return &mockAppointmentStore{
		err:         nil,
		exists:      false,
		appointment: exampleAppointment,
		appointments: []appointmentstore.Appointment{
			exampleAppointment,
		},
	}
}

func NewPutErrrorMockAppointmentStore() appointmentstore.AppointmentStore {
	return &mockAppointmentStore{
		err:         nil,
		exists:      true,
		putError:    true,
		appointment: exampleAppointment,
		appointments: []appointmentstore.Appointment{
			exampleAppointment,
		},
	}
}

func NewsSuccesMockAppointmentStore() appointmentstore.AppointmentStore {
	return &mockAppointmentStore{
		err:         nil,
		exists:      true,
		putError:    false,
		appointment: exampleAppointment,
		appointments: []appointmentstore.Appointment{
			exampleAppointment,
		},
	}
}

func (m *mockAppointmentStore) PutAppointment(appointment appointmentstore.Appointment) (appointmentstore.Appointment, error) {
	if m.putError {
		return appointmentstore.Appointment{}, errors.New("Error storing appointment")
	}
	return m.appointment, m.err
}

func (m *mockAppointmentStore) GetAppointmentOnUid(appointmentUID string) (appointmentstore.Appointment, bool, error) {
	return m.appointment, m.exists, m.err
}

func (m *mockAppointmentStore) GetAppointmentsOnUserUid(userUID string) ([]appointmentstore.Appointment, error) {
	return m.appointments, m.err
}

func (m *mockAppointmentStore) GetAppointmentsOnStatus(status appointmentstore.AppointmentStatus) ([]appointmentstore.Appointment, error) {
	return m.appointments, m.err
}

var exampleAppointment = appointmentstore.Appointment{
	AppointmentUID: "myAppointmentUid",
	UserUID:        "myUserUid",
	DateTime:       "myDateTime",
	Location:       "myLocation",
	Topic:          "myTopic",
	Status:         appointmentstore.AppointmentStatusRequested,
}
