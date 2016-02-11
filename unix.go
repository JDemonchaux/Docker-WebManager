package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"net"
	"io"
)

type UnixHandler struct {
	path string
}

func (h *UnixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("unix", h.path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()

	res, err := c.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer res.Body.Close()

	for i, r := range res.Header {
		for _, s := range r {
			w.Header().Add(i, s)
		}
	}

	if _, err := io.Copy(w, res.Body); err != nil {
		log.Println(err)
	}
}

func tcpHandler(e string) http.Handler {
	u, err := url.Parse(e)
	if err != nil {
		log.Fatal(err)
	}
	return httputil.NewSingleHostReverseProxy(u)
}

func unixHandler(e string) http.Handler {
	return &UnixHandler{e}
}

func handlerAPI(sock string) http.Handler {
	mux := http.NewServeMux()
	var h http.Handler

	if strings.Contains(sock, "http") {
		h = tcpHandler(sock)
	} else {
		h = unixHandler(sock)
	}
	mux.Handle("/", http.StripPrefix("", h))
	return mux
}

func unixSock(sock string) {
	handler := handlerAPI(sock)
	if err := http.ListenAndServe("127.0.0.1:1234", handler); err != nil {
		log.Fatal(err)
	}
}
