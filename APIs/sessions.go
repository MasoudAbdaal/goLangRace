package APIs

import (
	"encoding/json"
	"goLangRace/Utils"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("veryHardS3creeet!"))

const sessionName = "BusinessLogicFlaws_+_RaceConditions"

const sessionMaxAge = 1000 //seconds

type Auth struct {
	Username string `json:"UsernamE"`
	Password string `json:"PassworD"`
}

type User struct {
	GUID     string `json:"guid"`
	IsActive bool   `json:"isActive"`
	Balance  string `json:"balance"`
	Picture  string `json:"picture"`
	Age      int    `json:"age"`
	EyeColor string `json:"eyeColor"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credential Auth

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

				Utils.AdminSession(w, r, session)

			}

		}

	}

	if credential.Username == "Admin" && credential.Password == "hard@dminP@$$w0rd" {

		Utils.AdminSession(w, r, session)
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

		Utils.SaveSession(w, r, session)

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

		jsonData, err := os.ReadFile(".\\usersData.json")

		if err == nil {
			var users []User
			if err := json.Unmarshal(jsonData, &users); err == nil {
				if responseData, err := json.Marshal(users); err == nil {

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(responseData)
					return
				}
			}
		}

		http.Error(w, "Failed while doing some administration oprations!", http.StatusAlreadyReported)
	}

}
