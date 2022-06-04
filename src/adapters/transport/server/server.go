package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/temperature-converter/models"
	"github.com/mikejeuga/temperature-converter/src/domain"
	"net/http"
	"strconv"
)

type Server struct {
}

type ConversionHandler struct {
	degreeConverter domain.Converter
}

func NewServer() *http.Server {
	s := &Server{}

	converter := domain.NewConverter(domain.ConvertCtoF, domain.ConvertFtoC)
	conversionHandler := NewConversionHandler(converter)

	router := mux.NewRouter()

	router.HandleFunc("/", s.Home)
	router.HandleFunc("/to-fahrenheit/{temp}", conversionHandler.ConvertCtoF)
	router.HandleFunc("/to-celsius/{temp}", conversionHandler.ConvertFtoC)

	return &http.Server{
		Addr:    ":8069",
		Handler: router,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Healthy, the temperature seems perfect!")
}

func (h ConversionHandler) ConvertCtoF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp, err := strconv.ParseFloat(vars["temp"], 64)
	if err != nil {
		http.Error(w, "Error converting temperature", http.StatusInternalServerError)
	}

	f, err := h.degreeConverter.ConvertCtoF(models.Celsius(temp))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, f)

}

func (h ConversionHandler) ConvertFtoC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	temp, err := strconv.ParseFloat(vars["temp"], 64)
	if err != nil {
		http.Error(w, "Error converting temperature", http.StatusInternalServerError)
	}

	f, err := h.degreeConverter.ConvertFtoC(models.Fahrenheit(temp))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, f)
}

func NewConversionHandler(degreeConverter domain.Converter) *ConversionHandler {
	return &ConversionHandler{degreeConverter: degreeConverter}
}
