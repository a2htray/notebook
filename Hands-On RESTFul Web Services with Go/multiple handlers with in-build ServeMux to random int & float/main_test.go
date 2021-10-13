package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRandomFloat(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/randomFloat")
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

func TestRandomInt(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/randomInt")
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
