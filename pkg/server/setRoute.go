package server

import (
	"encoding/json"
	"flights/pkg/types"
	"io"
	"log"
	"net/http"
)

// GeneratesSubRoutesOfRoute this endpoint generates subroutes from a main route
// Entrada: POST [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
func GeneratesSubRoutesOfRoute(w http.ResponseWriter, r *http.Request) {

	var rest types.RestFul

	flights := types.Flights{}
	data, err := io.ReadAll(r.Body)
	err = json.Unmarshal(data, &flights)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rest.AddError(err)
		err = json.NewEncoder(w).Encode(rest)
		if err != nil {
			log.Printf("setFlight().json.NewEncoder(w).Encode(rest).Error: %v", err)
		}
		return
	}

	rest.Success(flights.GetSubRoutes())
	err = json.NewEncoder(w).Encode(rest)
	if err != nil {
		log.Printf("setFlight().json.NewEncoder(w).Encode(rest).Error: %v", err)
	}
}
