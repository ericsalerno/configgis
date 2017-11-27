package main

import "testing"

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
	s := NewServer(337, "something", 123)

	if s.getMatch.MatchString("/get/site.com/live/value") != true {
		t.Fatal("invalid get match regex!")
	}

	if s.setMatch.MatchString("/set/site.com/live/value") != true {
		t.Fatal("Invalid set match regex!")
	}
}
