package APIs

import (
	"encoding/json"
	en "goLangRace/Entities"
	"goLangRace/Utils"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("veryHardS3creeet!"))

const sessionName = "BusinessLogicFlaws_+_RaceConditions"

const sessionMaxAge = 1000 //seconds

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credential en.Auth

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := store.Get(r, sessionName)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if connString, ok := session.Values["Connection-String"].(string); ok {
		if _, ok := session.Values["Role"].(string); ok {

			if connString == "BANK_DB-ADMIN-CONNECTION_STRING" {

				Utils.AdminSession(&w, r, session)

			}

		}

	}

	if credential.Username == "Admin" && credential.Password == "hard@dminP@$$w0rd" {
		Utils.AdminSession(&w, r, session)
	}

	if credential.Username == "regular_user" && credential.Password == "password" {

		session, err := store.New(r, sessionName)
		if err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		session.Options.MaxAge = sessionMaxAge
		session.Values["Role"] = "Simple User"
		session.Values["Connection-String"] = "REGULAR_ConnectionString"

		Utils.SaveSession(&w, r, session)

	}

}
