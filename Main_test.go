package main

import (
	"encoding/json"
	"glutton/GStruct"
	"glutton/Handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler.HomeHandler)
	statusCheck(rr, t)

	handler.ServeHTTP(rr, req)

	recieved := &GStruct.RANDOMID{}
	err = json.Unmarshal(rr.Body.Bytes(), recieved)
	if err != nil {
		t.Errorf("Unable to parse JSON: Got %v", rr.Body.String())

	}

	if !(recieved.ID > 0 && recieved.ID < 9999) {
		t.Errorf("Id not in range: got %v want in range %v to %v",
			recieved.ID, 0, 9999)
	}

}

func statusCheck(rr *httptest.ResponseRecorder, t *testing.T) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
