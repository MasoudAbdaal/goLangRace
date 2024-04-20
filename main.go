package main

import (
	"goLangRace/APIs"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/scenario_1", APIs.TransferHandler)

	r.HandleFunc("/scenario_2/login", APIs.LoginHandler).Methods("POST")
	r.HandleFunc("/scenario_2/bank_balance", APIs.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/scenario_2/private", APIs.GetAllUsersHandler).Methods("GET")

	http.ListenAndServe(":8585", r)
}
