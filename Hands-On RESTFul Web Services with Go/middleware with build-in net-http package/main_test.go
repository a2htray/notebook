package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/api/book")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("return: %s", string(bytes))
}
