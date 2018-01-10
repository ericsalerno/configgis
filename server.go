package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-redis/redis"
)

// Server main object
type Server struct {
	port int

	redisDB *redis.Client

	getMatch *regexp.Regexp
	setMatch *regexp.Regexp
}

// NewServer creates a new server instance
func NewServer(port int, redisHost string, redisPort int) *Server {
	s := Server{}
	s.port = port

	connection := fmt.Sprintf("%s:%d", redisHost, redisPort)
	s.redisDB = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	s.getMatch = regexp.MustCompile("/get/([a-zA-Z0-9\\._-]+)/([a-zA-Z0-9\\._-]+)/([a-zA-Z0-9\\._-]+)/?")
	s.setMatch = regexp.MustCompile("/set/([a-zA-Z0-9\\._-]+)/([a-zA-Z0-9\\._-]+)/?")

	return &s
}

// AddConfigDirectory adds a configuration directory to load
func (s *Server) AddConfigDirectory(directory string) {

}

// Listen starts the server and listen for connections
func (s *Server) Listen() {
	localServer := fmt.Sprintf(":%d", s.port)
	http.ListenAndServe(localServer, s)
}

// ServeHTTP server HTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	getMatches := s.getMatch.FindSubmatch([]byte(r.URL.Path))

	if len(getMatches) == 4 {
		s.getValue(w, string(getMatches[1]), string(getMatches[2]), string(getMatches[3]))
		return
	}

	setMatches := s.setMatch.FindSubmatch([]byte(r.URL.Path))

	if len(setMatches) == 3 {
		r.ParseForm()
		if r.Method == "POST" {
			s.setValue(w, string(setMatches[1]), string(setMatches[2]), r.Form)
			return
		}
	}

	err := FailResponse()
	s.processResponse(err, w)
}

// setValue expects a url in the format of // /set/server/stage/key
func (s *Server) setValue(w http.ResponseWriter, server string, stage string, formValues url.Values) {

	successful := true
	count := 0
	for key, value := range formValues {
		val := s.redisDB.Set(key, value, 0)
		if val.Err != nil {
			successful = false
			break
		}
		count++
	}

	if successful {
		response := ValidResponse("Success", fmt.Sprintf("%d records updated.", count))
		s.processResponse(response, w)
		return
	}

	response := ErrorResponse(errors.New("An error occurred!"))
	s.processResponse(response, w)
}

func (s *Server) getValue(w http.ResponseWriter, server string, stage string, key string) {

	response := ValidResponse(key, "got some value")

	s.processResponse(response, w)
}

func (s *Server) processResponse(r Response, w http.ResponseWriter) {
	output, err := json.Marshal(r)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"status":500,"message":"Unable to process response."}`))
		return
	}

	w.WriteHeader(r.Status)
	w.Write(output)
}
