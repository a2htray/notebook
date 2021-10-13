package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/health")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("timestamp: %s", string(bytes))
}
