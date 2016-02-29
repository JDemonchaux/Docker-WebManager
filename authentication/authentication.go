package authentication

import (
	"net/http"
	"time"
	"crypto/rand"
	"log"
	"text/template"
)

type AuthenticationType struct {
	Username string
	Password string
	Token    string
}

func (auth *AuthenticationType)Auth(w http.ResponseWriter, req *http.Request) {
	expiration := time.Now().Add(time.Hour)
	cookie := &http.Cookie{Name: "athen",
		Value: auth.Token,
		HttpOnly: false,
		Expires: expiration,
		MaxAge: 3600}
	http.SetCookie(w, cookie)

	tmpl, err := template.ParseFiles("appWeb/header.html", "appWeb/login.html", "appWeb/footer.html")

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "index", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
	req.Body.Close()
}

func New(username string, password string) (*AuthenticationType) {
	var token = make([]byte, 2048)
	rand.Read(token)
	log.Println(string(token))
	return &AuthenticationType{username, password, string(token)}
}

func ValidAuth(pattern string, handler func(w http.ResponseWriter, req *http.Request)) {
	//DefaultServeMux.HandleFunc(pattern, handler)
}
