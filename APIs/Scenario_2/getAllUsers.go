package APIs

import (
	"encoding/json"
	en "goLangRace/Entities"
	"goLangRace/Utils"
	"net/http"
	"os"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	session, err := Utils.GetSession(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["UserID"].(string) == "0" {

		jsonData, err := os.ReadFile(".\\usersData.json")

		if err == nil {
			var users []en.User
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
		return
	} else {

		http.Error(w, "Get out of here!", http.StatusUnauthorized)
		return
	}

}
