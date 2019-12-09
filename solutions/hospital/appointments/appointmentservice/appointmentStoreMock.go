package main

type mockAppointmentStore struct {
	err          error
	exists       bool
	appointment  Appointment
	appointments []Appointment
}

func NewErrorMockAppointmentStore(err error) AppointmentStore {
	return &mockAppointmentStore{
		err:          err,
		appointment:  Appointment{},
		appointments: []Appointment{},
	}
}

func NewNotFoundMockAppointmentStore() AppointmentStore {
	return &mockAppointmentStore{
		err:         nil,
		exists:      false,
		appointment: exampleAppointment,
		appointments: []Appointment{
			exampleAppointment,
		},
	}
}

func NewsSuccesMockAppointmentStore() AppointmentStore {
	return &mockAppointmentStore{
		err:         nil,
		exists:      true,
		appointment: exampleAppointment,
		appointments: []Appointment{
			exampleAppointment,
		},
	}
}

func (m *mockAppointmentStore) PutAppointment(appointment Appointment) (Appointment, error) {
	return m.appointment, m.err
}

func (m *mockAppointmentStore) GetAppointmentOnUid(appointmentUID string) (Appointment, bool, error) {
	return m.appointment, m.exists, m.err
}

func (m *mockAppointmentStore) GetAppointmentsOnUserUid(userUID string) ([]Appointment, error) {
	return m.appointments, m.err
}

func (m *mockAppointmentStore) GetAppointmentsOnStatus(status AppointmentStatus) ([]Appointment, error) {
	return m.appointments, m.err
}

var exampleAppointment = Appointment{
	AppointmentUID: "myAppointmentUid",
	UserUID:        "myUserUid",
	DateTime:       "myDateTime",
	Location:       "myLocation",
	Topic:          "myTopic",
	Status:         AppointmentStatusRequested,
}
