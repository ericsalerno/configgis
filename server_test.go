package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer(337, "something", 123)

	if s.port != 337 {
		t.Fatal("Invalid port found!")
	}

	if s.redisDB == nil {
		t.Fatal("Invalid redis db client!")
	}

	if s.getMatch == nil {
		t.Fatal("Get match invalid!")
	}

	if s.setMatch == nil {
		t.Fatal("Set match invalid!")
	}
}

func TestRoutingMatches(t *testing.T) {
	s := NewServer(337, "localhost", 6379)

	if s.getMatch.MatchString("/get/site.com/live/value") != true {
		t.Fatal("invalid get match regex!")
	}

	if s.setMatch.MatchString("/set/site.com/live/value") != true {
		t.Fatal("Invalid set match regex!")
	}
}

func TestSetValue(t *testing.T) {
	s := NewServer(123, "localhost", 6379)

	w := httptest.NewRecorder()

	s.setValue(w, "something", "live", "key")

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Fatal("Invalid status code received!")
	}

	bodyStr := string(body)

	fmt.Println(bodyStr)
	if strings.Contains(bodyStr, `"key":"key"`) == false {
		t.Fatal("Did not find key field!")
	}
}

func TestServeHTTP(t *testing.T) {
	s := NewServer(123, "localhost", 6379)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/set/seomthing/something/something", nil)

	s.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Fatal("Invalid status code received!")
	}

	bodyStr := string(body)

	fmt.Println(bodyStr)
	if strings.Contains(bodyStr, `"key":"key"`) == false {
		t.Fatal("Did not find key field!")
	}
}
