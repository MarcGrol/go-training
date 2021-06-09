install-swagger:
	sh install_go_swagger.sh

generate:
	mkdir -p service
	cd service/ && find . ! -name \configure_*.go -type f -exec rm -f {} +
	cd service/ && find . -type d -empty -delete
	cd service/ && ../bin/swagger generate server -A patient-service -f ../contract.yml

run:
	go run service/cmd/patient-service-server/main.go --port 8080

.PHONY: install-swagger generate