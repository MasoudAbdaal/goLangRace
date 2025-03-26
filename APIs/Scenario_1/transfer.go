package APIs

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var accountBalance int = 100

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("n")

	inputNum, err := strconv.Atoi(amount)

	if err != nil {
		http.Error(w, "please check ?n= parameter!", http.StatusBadRequest)
		// return
	}

	if inputNum != 1 {
		http.Error(w, "Your input should be 1", http.StatusBadRequest)
		// return
	}

	fmt.Printf("accountBalance: %d\n", accountBalance)

	//Sub-state
	accountBalance += inputNum

	fmt.Fprintf(w, "User Account Balance: %d \n", accountBalance)

	//Race window!
	time.Sleep(2 * time.Millisecond) //Submit changes to DB
	//-------

	accountBalance = 100 //Example to integrity check function operations!

	body, err := io.ReadAll(r.Body)
	if err == nil {
		fmt.Fprintf(w, "Recieved Request Body:\n %s", body)
	}
	defer r.Body.Close()
	// test signed commit
	// test signed commit
	// test signed commit
	// test signed commit
}
