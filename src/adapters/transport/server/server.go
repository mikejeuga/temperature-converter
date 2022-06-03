package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Healthy, the temperature seems perfect!")
}

func NewServer() *http.Server {
	s := Server{}

	router := mux.NewRouter()

	router.HandleFunc("/", s.Home)
	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
