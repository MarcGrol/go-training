#!/bin/sh

# create
export PATIENT_UID=$(curl -s \
    -X POST \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"fullName":"Marc Grol","addressLine":"Heemstrakwartier 79","allergies":["gezeik"]}' \
     https://gotrainingxebia.appspot.com/api/patient | jq --raw-output '.uid')

# get
curl -s \
    -X GET \
    -H 'Accept: application/json' \
    https://gotrainingxebia.appspot.com/api/patient/${PATIENT_UID}

# modify
curl -s \
    -X PUT \
    -H 'Accept: application/json' -H 'Content-Type: application/json' \
     --data '{"fullName":"Marc Grol","addressLine":"Thuis","allergies":["gedoe"]}' \
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
