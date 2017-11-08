package configgis

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	//s.redisHost = redisHost
	//s.redisPort = redisPort
	connection := fmt.Sprintf("%s:%d", redisHost, redisPort)
	s.redisDB = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	s.getMatch = regexp.MustCompile("/get/(.*)")
	s.setMatch = regexp.MustCompile("/set/(.*)")

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
	if s.setMatch.MatchString(r.URL.Path) {
		s.setValue(w, r)
	} else if s.getMatch.MatchString(r.URL.Path) {
		s.getValue(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (s *Server) setValue(w http.ResponseWriter, r *http.Request) {

	response := ValidResponse("set", r.URL.Path)

	s.processResponse(response, w)
}

func (s *Server) getValue(w http.ResponseWriter, r *http.Request) {

	response := ValidResponse("get", r.URL.Path)

	s.processResponse(response, w)
}

func (s *Server) processResponse(r Response, w http.ResponseWriter) {
	output, err := json.Marshal(r)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"status\":500,\"message\":\"Unable to process response.\"}"))
		return
	}

	w.Write(output)
}
