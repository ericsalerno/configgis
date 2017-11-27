package main

import (
	"errors"
	"testing"
)

func TestValidResponse(t *testing.T) {
	r := ValidResponse("something", "value")

	if r.Key != "something" {
		t.Fatal("Invalid key, got " + r.Key + " was expecting 'something'")
	}

	if r.Message != "OK" {
		t.Fatal("Invalid message, got " + r.Message + " was expecting 'OK'")
	}

	if r.Status != 200 {
		t.Fatal("Invalid status code, was expecting 200")
	}

	if r.Value != "value" {
		t.Fatal("Invalid value, was expecting 'value' got " + r.Value)
	}
}

func TestErrorResponse(t *testing.T) {
	err := errors.New("Oh noes!")
	r := ErrorResponse(err)

	if r.Key != "" {
		t.Fatal("Invalid key, got " + r.Key + " was expecting ''")
	}

	if r.Message != "Oh noes!" {
		t.Fatal("Invalid message, got '" + r.Message + "' was expecting 'Oh noes!'")
	}

	if r.Status != 500 {
		t.Fatal("Invalid status code, was expecting 500")
	}

	if r.Value != "" {
		t.Fatal("Invalid value, was expecting '' got '" + r.Value + "'")
	}
}

func TestFailResponse(t *testing.T) {
	r := FailResponse()

	if r.Key != "" {
		t.Fatal("Invalid key, got " + r.Key + " was expecting ''")
	}

	if r.Message == "" {
		t.Fatal("Invalid message, got '" + r.Message + "' was expecting something.")
	}

	if r.Status != 400 {
		t.Fatal("Invalid status code, was expecting 500")
	}

	if r.Value != "" {
		t.Fatal("Invalid value, was expecting '' got '" + r.Value + "'")
	}
}
