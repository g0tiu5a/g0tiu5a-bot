package common

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const (
	BUFSIZE   = 1024
	TEST_FILE = "./test_data/event_1.json"
)

func TestDecode(t *testing.T) {
	// Open json_file for testing
	buf, err := ioutil.ReadFile(TEST_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// Create HTTP Response
	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(buf)),
	}

	// Decode
	var events []interface{}
	Decode(response, &events)
}
