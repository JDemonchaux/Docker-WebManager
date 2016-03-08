package authentication

import (
	"time"
	"crypto/rand"
	"log"
	"text/template"
	"net/http"
	"crypto/sha256"
	"io"
	"encoding/hex"
	"fmt"
)

type AuthenticationType struct {
	Username      string
	Password      string
	Token         string
	Authenticated bool
	Date          time.Time
}


func (auth *AuthenticationType)IsAuthenticated(req *http.Request) bool {
	if auth.Authenticated{
		return true
	}
	cookie, err := req.Cookie("auth")
	if err != nil{
		return false
	}

	log.Println(auth.Token)
	log.Println("-------------------------------------------------------------------------------------------------")
	log.Println(cookie.Value)
	return true
}

func (auth *AuthenticationType)Auth(w http.ResponseWriter, req *http.Request) {


	authHash := req.FormValue("auth");

	h256 := sha256.New()
	io.WriteString(h256, auth.Username + auth.Password)
	hash := hex.EncodeToString(h256.Sum(nil))

	if authHash == string(hash) {
		log.Println("Authenticated")
		auth.Authenticated = true
	}

	expiration := time.Now().Add(time.Hour)
	cookie := &http.Cookie{Name: "auth",
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
	return &AuthenticationType{username, password, fmt.Sprintf("%x", token), false, time.Now()}
}



