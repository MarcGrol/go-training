package registrationService

type Patient struct {
	UID         string
	Name        string
	PhoneNumber string
}

//go:generate mockgen -source=externalApis.go -destination=patientServiceMocks.go -package=registrationService PatientStorer,PincodeGenerator,SmsSender

// START OMIT
type PatientStorer interface {
	GetPatientOnUid() (Patient, error)
	PutPatientOnUid(patient Patient) error
}

type PincodeGenerator interface {
	GeneratePincode() int
}

type SmsSender interface {
	SendSms(phoneNumber string, smsContent string) error
}

// END OMIT
