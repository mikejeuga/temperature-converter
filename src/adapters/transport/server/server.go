package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/temperature-converter/src/domain"
	"net/http"
)

type Server struct {
}

func NewServer() *http.Server {
	s := &Server{}

	conversionHandler := NewConversionHandler(domain.Converter(domain.ConvertCtoF))

	router := mux.NewRouter()

	router.HandleFunc("/", s.Home)
	router.HandleFunc("/fahrenheit/{temp}", conversionHandler.ConvertCtoF)
	return &http.Server{
		Addr:    ":8069",
		Handler: router,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Healthy, the temperature seems perfect!")
}

type ConversionHandler struct {
	degreeConverter domain.Converter
}

func (h ConversionHandler) ConvertCtoF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}

func NewConversionHandler(degreeConverter domain.Converter) *ConversionHandler {
	return &ConversionHandler{degreeConverter: degreeConverter}
}
