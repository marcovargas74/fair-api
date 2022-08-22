package entity

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"
)

const (
	erroMsg       = "Test fail got Value[%v], wait Value [%v]"
	UserAgentTest = "self_test"
)

//CheckIfEqualString check if result is OK type string
func CheckIfEqualString(t *testing.T, gotValue, waitValue string) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualBool check if result is OK type BOOL
func CheckIfEqualBool(t *testing.T, gotValue, waitValue bool) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualInt check if result is OK type INT
func CheckIfEqualInt(t *testing.T, gotValue, waitValue uint64) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualFloat check if result is OK type FLOAT
func CheckIfEqualFloat(t *testing.T, gotValue, waitValue float64) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualTime check if result is OK type Time
func CheckIfEqualTime(t *testing.T, gotValue, waitValue time.Time) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfUptimeIsOK check if result is OK to UpTime
func CheckIfUptimeIsOK(t *testing.T, gotValue, waitValue float64) {
	t.Helper()
	if gotValue <= waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//NewReqEndpointsGET Genetic GET endpoint to test
func NewReqEndpointsGET(urlPrefix, urlName string) *http.Request {
	request, error := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)
	if error != nil {
		panic(error)
	}

	request.Header.Set("User-Agent", UserAgentTest)
	fmt.Printf("endpoint: %v\n", request.URL)
	return request
}

//NewReqEndpointsPOST Genetic POST endpoint to test
func NewReqEndpointsPOST(urlPrefix, urlName string) *http.Request {
	request, error := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", urlPrefix, urlName), nil)
	if error != nil {
		panic(error)
	}

	request.Header.Set("User-Agent", UserAgentTest)
	return request
}

//NewReqEndpointsBodyPOST Genetic POST endpoint to test
func NewReqEndpointsBodyPOST(urlPrefix, urlName string, json string) *http.Request {

	//jsonBody := []byte(`{"client_message": "hello, server!"}`)
	jsonBody := []byte(`{
		"name": "VILA TESTE",
		"District": "dstrito teste",
		"Region5": "test",
		"neighborhood" :"bairo"
	}
      `)

	bodyReader := bytes.NewReader(jsonBody)

	request, error := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", urlPrefix, urlName), bodyReader)
	if error != nil {
		panic(error)
	}

	request.Header.Set("User-Agent", UserAgentTest)
	return request
}
