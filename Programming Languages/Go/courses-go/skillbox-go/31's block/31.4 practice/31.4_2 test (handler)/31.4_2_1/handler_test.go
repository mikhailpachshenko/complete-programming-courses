package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handler))

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()

	text := string(textBytes)
	if text != "message from test handler" {
		t.Log(text)
		t.Fail()
	}
}
