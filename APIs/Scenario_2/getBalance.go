package APIs

import (
	"encoding/json"
	"fmt"
	"goLangRace/Utils"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	session, err := Utils.GetSession(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if _, ok := session.Values["UserID"].(string); !ok {

		http.Error(w, "Something went wrong while getting your sessions info...Login again", http.StatusUnauthorized)
		return

	}

	//Programming mistake lead to race! --Race Window--
	session.Values["Connection-String"] = "BANK_DB-ADMIN-CONNECTION_STRING"

	var userBalance int = GetBalanceFromBank(session)

	s, _ := Utils.GetSession(r)
	fmt.Println(s.Values["Connection-String"])

	session.Values["Connection-String"] = "REGULAR_ConnectionString"
	fmt.Println(s.Values["Connection-String"])

	//--------------Sub-state ends here-------------------

	responseData := map[string]interface{}{
		"Balance": userBalance,
		"UserID":  session.Values["UserID"],
	}

	jsonData, err := json.Marshal(responseData)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}

}

func GetBalanceFromBank(ConnectionInfo *sessions.Session) int {
	fmt.Println("Connect To Bank DB Using Connections Info....")
	time.Sleep(7000 * time.Millisecond)
	fmt.Printf("Result is 8585 $ for userId %s\n", (ConnectionInfo.Values["UserID"].(string)))
	return 8585
}
