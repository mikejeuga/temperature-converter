package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/temperature-converter/models"
	"github.com/mikejeuga/temperature-converter/src/internal/domain"
	"net/http"
	"strconv"
)

type ConversionHandler struct {
	degreeConverter domain.Converter
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
