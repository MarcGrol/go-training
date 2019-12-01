

curl -X GET  -H 'Accept:application/json' \
     http://localhost:8080/api/patient/123

curl -X POST -H 'Content-type:application/json' -H 'Accept:application/json' \
     --data '{"FullName":"Marc","AddressLine":"Heemstra","Allergies":["pinda"]}' \
     http://localhost:8080/api/patient

