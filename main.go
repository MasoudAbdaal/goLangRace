package main

import (
	Scenario_1 "goLangRace/APIs/Scenario_1"
	Scenario_2 "goLangRace/APIs/Scenario_2"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/scenario_1", Scenario_1.TransferHandler)

	r.HandleFunc("/scenario_2/login", Scenario_2.LoginHandler).Methods("POST")
	r.HandleFunc("/scenario_2/bank_balance", Scenario_2.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/scenario_2/private", Scenario_2.GetAllUsersHandler).Methods("GET")

	http.ListenAndServe(":8585", r)
}
