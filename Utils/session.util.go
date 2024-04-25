package Utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "BusinessLogicFlaws_+_RaceConditions"

var store = sessions.NewCookieStore([]byte("veryHardS3creeet!"))

const sessionMaxAge = 1000 //seconds

func SaveSession(w *http.ResponseWriter, r *http.Request, session *sessions.Session) {

	err := session.Save(r, *w)
	if err != nil {
		http.Error(*w, err.Error()+"\n Try again....", http.StatusInternalServerError)
		return
	}

}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}

func CreateSession(r *http.Request, w *http.ResponseWriter) {
	session, err := store.New(r, sessionName)
	if err != nil {

		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}
	session.Options.MaxAge = sessionMaxAge
	session.Values["UserID"] = "1"
	session.Values["Connection-String"] = "REGULAR_ConnectionString"

	SaveSession(w, r, session)

}

func AdminSession(w *http.ResponseWriter, r *http.Request, session *sessions.Session) {

	session.Options.MaxAge += 9999
	session.Values["UserID"] = "0"
	session.Options.Domain = "localhost:8085, private.localhost:8085"

	SaveSession(w, r, session)

	http.Redirect(*w, r, "/scenario_2/private", http.StatusPermanentRedirect)
}
