package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/temperature-converter/src/adapters/transport/server/handlers"
	"github.com/mikejeuga/temperature-converter/src/domain"
	"net/http"
)

type Server struct {
	temperatureHandler handlers.ConversionHandler
}

func NewServer() *http.Server {

	converter := domain.NewConverter(domain.ConvertCtoF, domain.ConvertFtoC)
	conversionHandler := handlers.NewConversionHandler(converter)
	s := &Server{temperatureHandler: *conversionHandler}

	router := mux.NewRouter()

	router.HandleFunc("/", s.Home)
	router.HandleFunc("/to-fahrenheit/{temp}", s.temperatureHandler.ConvertCtoF)
	router.HandleFunc("/to-celsius/{temp}", s.temperatureHandler.ConvertFtoC)

	return &http.Server{
		Addr:    ":8069",
		Handler: router,
	}
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Healthy, the temperature seems perfect!")
}
