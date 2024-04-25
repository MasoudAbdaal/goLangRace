package APIs

import (
	"encoding/json"
	"fmt"
	en "goLangRace/Entities"
	"goLangRace/Utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credential en.Auth
	session, err := Utils.GetSession(r)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Login info Session:=> %s \n", session.Values["Connection-String"])

	if connString, ok := session.Values["Connection-String"].(string); ok {
		if _, ok := session.Values["UserID"].(string); ok {

			if connString == "BANK_DB-ADMIN-CONNECTION_STRING" {

				Utils.AdminSession(&w, r, session)

			}

		}

	}
	err = json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if credential.Username == "Admin" && credential.Password == "hard@dminP@$$w0rd" {
		Utils.AdminSession(&w, r, session)
	}

	if credential.Username == "regular_user" && credential.Password == "password" {

		Utils.CreateSession(r, &w)

	}

}
