package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGoVersion(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/api/v1/go-version")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("random float: %s", string(bytes))
}

func TestGetFileContent(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/api/v1/show-file/foo")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("random int: %s", string(bytes))
}
