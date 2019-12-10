#!/bin/sh

# create
export PATIENT_UID=$(curl -s \
    -X POST \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"FullName":"Marc Grol","AddressLine":"Heemstrakwartier 79","Allergies":["gezeik"]}' \
     https://gotrainingxebia.appspot.com/api/patient | jq --raw-output '.UID')

# get
curl -s \
    -X GET \
    -H 'Accept: application/json' \
    https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

# modify
curl -s \
    -X PUT \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"FullName":"Marc Grol","AddressLine":"Thuis","Allergies":["gedoe"]}' \
     https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

# get
curl -s \
    -X GET \
    -H 'Accept: application/json' \
    https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

# delete
curl -s \
    -X DELETE \
     https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

# get
curl -s \
    -X GET \
    -H 'Accept: application/json' \
    https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}
