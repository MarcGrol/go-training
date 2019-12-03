package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Patient struct {
	Name string
	Age int
}

type PatientClient struct {
	HostName string
}

func (pc PatientClient) Fetch(uid string) (Patient,error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/api/patient/%s", pc.HostName, uid))
	if err != nil {
		return Patient{}, fmt.Errorf("Error fetching patient: %s", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != 200 {
		return Patient{}, fmt.Errorf("Error fetching patient: http-status %d", httpResponse.StatusCode)
	}
	dec := json.NewDecoder(httpResponse.Body)
	var resp Patient
	err = dec.Decode(&resp)
	if err != nil {
		return Patient{}, fmt.Errorf("Error decoding json response: %s", err)
	}
	return resp, nil
}