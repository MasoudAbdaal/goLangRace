package Utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func SaveSession(w *http.ResponseWriter, r *http.Request, session *sessions.Session) {

	err := session.Save(r, *w)
	if err != nil {
		http.Error(*w, err.Error()+"\n Try again....", http.StatusInternalServerError)
		return
	}

}

func AdminSession(w *http.ResponseWriter, r *http.Request, session *sessions.Session) {

	session.Options.MaxAge += 9999
	session.Values["Role"] = "Admin"
	session.Options.Domain = "localhost:8085, private.localhost:8085"

	SaveSession(w, r, session)

	http.Redirect(*w, r, "/scenario_2/private", http.StatusPermanentRedirect)
}
