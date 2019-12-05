package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s mode", os.Args[0])
	}
	mode := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if mode == "internal" {
		client, cleanup, err := appointmentapi.NewExternalGrpcClient(appointmentapi.DefaultPort)
		if err != nil {
			log.Fatalf("*** Error creating external-appointment-client: %v", err)
		}
		defer cleanup()

		{
			resp, err := client.GetAppointmentsOnUser(ctx, &appointmentapi.GetAppointmentsOnUserRequest{
				UserUid: "1",
			})
			if err != nil {
				log.Fatalf("getting appointments on user: %s", err)
			}
			if resp.Error != nil {
				log.Fatalf("Error getting appointments on user: %+v", resp.Error)
			}
			log.Printf("appointments for user:%+v", resp.Appointments)
		}

		{
			resp, err := client.RequestAppointment(ctx, &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  "1",
					DateTime: "now",
					Location: "Leuven",
					Topic:    "Onderzoek",
				},
			})
			if err != nil {
				log.Fatalf("request appointment: %s", err)
			}
			if resp.Error != nil {
				log.Fatalf("Error requesting appointment: %+v", resp.Error)
			}
			log.Printf("requested appointment:%+v", resp.Appointment)
		}

	} else if mode == "external" {
		client, cleanup, err := appointmentapi.NewInternalGrpcClient(appointmentapi.DefaultPort)
		if err != nil {
			log.Fatalf("*** Error creating external-appointment-client: %v", err)
		}
		defer cleanup()
		{
			resp, err := client.GetAppointmentsOnStatus(ctx, &appointmentapi.GetAppointmentsOnStatusRequest{
				Status: appointmentapi.AppointmentStatus_REQUESTED,
			})
			if err != nil {
				log.Fatalf("getting appointments on user: %s", err)
			}
			if resp.Error != nil {
				log.Fatalf("Error getting appointments: %+v", resp.Error)
			}
			log.Printf("appointments with status: %+v", resp.Appointments)
		}

		{
			resp, err := client.ModifyAppointmentStatus(ctx, &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "a",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			})
			if err != nil {
				log.Fatalf("confirm appointment: %s", err)
			}
			if resp.Error != nil {
				log.Fatalf("Error confirming appointment: %+v", resp.Error)
			}
			log.Printf("confirmed appointment:%+v", resp.Appointment)
		}
	} else {
		log.Fatalf("Unknown mode %s", mode)
	}
}
