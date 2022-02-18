# GRPC Registration client

This will connect to a GRPC server at localhost on port 60002

## Compile the client

    go build

This will result in client executable in the current directory

## Starting the registration of a patient

    ./client -command=start-registration -bsn=9999999 -name=marc -email=mgrol@xebia.com

This will return a patient-uid and will result in an email with pin-code being sent to mgrol@xebia.com

## Completing registration of a patient

Use the patient-uid from above and pin-code from the mail to start this command

    ./client -command=complete-registration -patient-uid=<patient-uid> -pin-code=<pin-code>

## Brute-force guessing the pin-code to complete the registration of a patient

Use the patient-uid from above to start this command

    ./client -command=bruteforce-registration -patient-uid=c2a755a3-78b1-423f-8794-7da967af52cd

