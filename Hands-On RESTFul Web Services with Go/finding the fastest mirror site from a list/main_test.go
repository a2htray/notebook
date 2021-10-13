package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	if resp, err := http.Get("http://localhost:8000/fastest-mirror"); err != nil {
		t.Fatal(err)
	} else {
		defer resp.Body.Close()

		t.Logf("response status: %s\n", resp.Status)

		respJSON, _ := ioutil.ReadAll(resp.Body)
		t.Logf("JSON string: %s\n", string(respJSON))
	}
}
