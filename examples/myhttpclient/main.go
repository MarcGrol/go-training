package myhttpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	hostname string
	timeout  time.Duration
}

func New(hostname string) *Client {
	return &Client{
		hostname: hostname,
		timeout:  5 * time.Second,
	}
}

// START OMIT
type GetPatientResponse struct {
	UID       string   `json:"uid"`
	FullName  string   `json:"fullName"`
	Allergies []string `json:"allergies"`
}

func (cl *Client) GetClientOnUid(patientUid string) (*GetPatientResponse, error) {
	client := http.Client{Timeout: cl.timeout}
	httpResponse, err := client.Get(fmt.Sprintf("http://abc.com/api/patient/", patientUid))
	if err != nil {
		return nil, fmt.Errorf("Error fetching patient: %s", err)
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("Error fetching patient: http-status %d", httpResponse.StatusCode)
	}
	dec := json.NewDecoder(httpResponse.Body)
	var resp GetPatientResponse
	err = dec.Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("Error decoding json response: %s", err)
	}
	return &resp, nil
}

// END OMIT
