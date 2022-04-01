package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestSendRequest(t *testing.T) {
  normalServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
  defer normalServer.Close()
  InitLog()

  err := sendRequest(normalServer.URL)
  if err != nil {
    t.Errorf("get wrong with request sending.")
  }

  abnormalServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusGatewayTimeout)
	}))
  defer abnormalServer.Close()
  err = sendRequest(abnormalServer.URL)
  if err == nil {
    t.Errorf("get wrong with request retrying.")
  }
}
