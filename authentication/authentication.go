package authentication

import (
	"net/http"
	"io"
	"time"
	"crypto/rand"
)

type AuthenticationType struct {
	Username string
	Password string
}

var Token = make([]byte, 2048)

func (AuthenticationType)Auth(w http.ResponseWriter, req *http.Request) {
	rand.Read(Token)

	expiration := time.Now().Add(time.Hour)
	cookie := &http.Cookie{Name: "athen",
		Value: string(Token),
		HttpOnly: false,
		Expires: expiration,
		MaxAge: 3600}
	http.SetCookie(w, cookie)
	io.WriteString(w, "Hello world!")
}
