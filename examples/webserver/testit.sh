#!/bin/sh

export PATIENT_UID="2ef04552-8983-48b2-b267-15aa48243d23"

curl -vvv \
    -X GET \
    -H 'Accept: application/json' \
    https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

curl -vvv \
    -X POST \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"FullName":"Marc Grol","AddressLine":"Heemstrakwartier 79","Allergies":["gezeik"]}' \
     https://gotrainingxebia.appspot.com/api/patient

curl -vvv \
    -X PUT \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"FullName":"Marc Grol","AddressLine":"Thuis","Allergies":["gedoe"]}' \
     https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

curl -vvv \
    -X DELETE \
     https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}