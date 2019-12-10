package appointmentstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOnUidFound(t *testing.T) {
	c := context.TODO()
	sut := NewAppointmentStore(NewBasicUuider())
	appointment, found, err := sut.GetAppointmentOnUid(c, "a")
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "a", appointment.AppointmentUID)
}

func TestGetOnUidNotFound(t *testing.T) {
	c := context.TODO()
	sut := NewAppointmentStore(NewBasicUuider())
	_, found, err := sut.GetAppointmentOnUid(c, "c")
	assert.NoError(t, err)
	assert.False(t, found)
}

func TestGetOnUserUid(t *testing.T) {
	c := context.TODO()
	sut := NewAppointmentStore(NewBasicUuider())
	appointments, err := sut.GetAppointmentsOnUserUid(c, "1")
	assert.NoError(t, err)
	assert.Len(t, appointments, 1)
	assert.Equal(t, "a", appointments[0].AppointmentUID)
	assert.Equal(t, "1", appointments[0].UserUID)
}

func TestGetOnStatus(t *testing.T) {
	c := context.TODO()
	sut := NewAppointmentStore(NewBasicUuider())
	appointments, err := sut.GetAppointmentsOnStatus(c, AppointmentStatusRequested)
	assert.NoError(t, err)
	assert.Len(t, appointments, 2)
	assert.Equal(t, AppointmentStatusRequested, appointments[0].Status)
	assert.Equal(t, AppointmentStatusRequested, appointments[1].Status)
}

func TestPut(t *testing.T) {
	c := context.TODO()
	sut := NewAppointmentStore(NewBasicUuider())
	appointment, err := sut.PutAppointment(c, Appointment{
		AppointmentUID: "c",
		UserUID:        "3",
		DateTime:       "now",
		Location:       "here",
		Topic:          "why",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, appointment.UserUID)

	newlyCreated, found, err := sut.GetAppointmentOnUid(c, appointment.AppointmentUID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.NotEmpty(t, newlyCreated.AppointmentUID)
	assert.Equal(t, "3", newlyCreated.UserUID)
	assert.Equal(t, "now", newlyCreated.DateTime)
	assert.Equal(t, "here", newlyCreated.Location)
	assert.Equal(t, "why", newlyCreated.Topic)
}
