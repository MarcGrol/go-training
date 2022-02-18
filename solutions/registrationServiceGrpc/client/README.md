# GRPC Appointment client

## Starting the client

    go install && appointmentclient external (patient asking for an appointment)
    go install && appointmentclient internal (nurse confirming appointment)

This will connect to a GRPC server at localhost on port 60001