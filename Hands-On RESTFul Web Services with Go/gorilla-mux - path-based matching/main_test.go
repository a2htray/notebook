package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/articles/books/123")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("detail: %s", string(bytes))
}
