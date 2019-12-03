package smsClient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	debug = true
	timeout = 10 * time.Second
	originatingNumber = "+12016043948"
	twilioURL =  "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"
)

type smsClient struct {
	accountId string
	password string
}

func NewTwilloSmsClient(accountId, password string) SmsClient {
	return &smsClient{
		accountId:accountId,
		password:password,
	}
}

func (sn *smsClient) SendSms(c context.Context, destinationNumber string, msgPayload string) error {
	client := &http.Client{
		Timeout:timeout,
	}

	form := url.Values{}
	form.Set("From", originatingNumber)
	form.Set("To", destinationNumber)
	form.Set("Body", msgPayload)

	req, err := http.NewRequest("POST", fmt.Sprintf(twilioURL, sn.accountId), strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("[twilio] create request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	req.SetBasicAuth(sn.accountId, sn.password)

	if debug {
		reqDump, err := httputil.DumpRequest(req, true)
		if err == nil {
			unescapedRequest, _ := url.QueryUnescape(string(reqDump))
			fmt.Printf( "[twilio] http request:\n %s", unescapedRequest)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf( "[twilio] send/receive error: %s", err.Error())
	}
	defer resp.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			fmt.Printf("[twilio] http response:\n%s", respDump)
		}
	}

	if resp.StatusCode >= 300 {
		// parse error response
		var response twilioErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return fmt.Errorf( "[twilio] decode response: %s", err)
		}

		if resp.StatusCode == 400 {
			// Be able to tell the ui that the number was invalid
			return fmt.Errorf("[twilio] Error sending sms %d: %+v", resp.StatusCode, response)
		}

		if resp.StatusCode > 400 {
			// All other errors are marked as internal errors
			return fmt.Errorf("[twilio] http error status %d, msg '%s', error-msg '%s'", resp.StatusCode, response.Message, response.ErrorMessage)
		}
	}
	
	// parse success response
	var response twilioSuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf( "[twilio] decode response: %s", err)
	}
	fmt.Printf( "[twilio] Successfully sent sms %d: %+v", resp.StatusCode, response)
	return nil
}

// @JsonStruct()
type twilioErrorResponse struct {
	// Error related fields
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message"`
	MoreInfo     string `json:"more_info"`
	Status       int    `json:"status,omitempty"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	URI          string `json:"uri"`
}

// @JsonStruct()
type twilioSuccessResponse struct {
	// Success related fields
	Sid         string `json:"sid"`
	AccountSid  string `json:"account_sid"`
	To          string `json:"to"`
	From        string `json:"from"`
	Body        string `json:"body"`
	Status      string `json:"status"`
	NumSegments string `json:"num_segments"`
	NumMedia    string `json:"num_media"`
	Direction   string `json:"direction"`
}
