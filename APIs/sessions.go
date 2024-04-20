package APIs

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("veryHardS3creeet!"))

const sessionName = "BusinessLogicFlaws_+_RaceConditions"

const sessionMaxAge = 1000 //seconds

type User struct {
	Username string `json:"UsernamE"`
	Password string `json:"PassworD"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credential User

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, _ := store.Get(r, sessionName)

	if (session.Values["Connection-String"].(string) == "BANK_DB-ADMIN-CONNECTION_STRING") ||
		(credential.Username == "Admin" && credential.Password == "hard@dminP@$$w0rd") {

		session.Options.MaxAge = session.Options.MaxAge + 99999
		session.Values["Role"] = "Admin"

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error()+"\n Try again....", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/scenario_2/private", http.StatusPermanentRedirect)
		return
	}

	if credential.Username == "regular_user" && credential.Password == "password" {

		session, err := store.New(r, "User-Session")

		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		session.Options.MaxAge = sessionMaxAge
		session.Values["Role"] = "Simple User"
		session.Values["Connection-String"] = "REGULAR_ConnectionString"

	}

}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	//if (!session)
	// 401 un-authorized()

	//-----RACE WINDOWS------
	//set (connection_string to BankDB_ADMIN!)
	// DB Operations
	//---------------------

	//Restore user session_connection_string to REGULAR

	//return data

}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, sessionName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["Role"].(string) == "Admin" {

		user := User{
			Username: "User1",
			Password: "C0mpl3xP@$$2-dd0",
		}

		jsonData, err := (json.Marshal(user))

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			return
		}

		http.Error(w, "Failed while doing some administration oprations!", http.StatusAlreadyReported)
	}

}
