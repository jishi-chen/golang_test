package Utility

import (
	"net/http"
	"text/template"
	"time"
)

func SetCookie(w http.ResponseWriter) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	//cookie, _ := r.Cookie("username")
	for _, cookie := range r.Cookies() {
		template.HTMLEscape(w, []byte(cookie.Name))
	}
}
