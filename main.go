package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8585", nil)
}

var accountBalance int = 100

func rootHandler(w http.ResponseWriter, r *http.Request) {
	inputQuery := r.URL.Query().Get("n")
	inputNum, err := strconv.Atoi(inputQuery)

	if inputNum != 1 {
		http.Error(w, "Your input should be 1", http.StatusBadRequest)
		return
	}

	if err == nil {
		fmt.Printf("accountBalance: %d\n", accountBalance)
	}

	accountBalance += inputNum //Account transfer operations
	fmt.Fprintf(w, "User Account Balance: %d \n", accountBalance)
	time.Sleep(2 * time.Millisecond) //Submit to DB
	accountBalance = 100             //Example to integrity check function operations!

}
func hello() {
	fmt.Println("HELLO WORLDS")
}
