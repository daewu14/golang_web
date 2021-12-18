package belajar_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{}
	cookie.Name = "x-daw-name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success set cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("x-daw-name")
	if err != nil {
		fmt.Fprint(w,"No cookie")
	} else {
		// name := cookie.Name
		value := cookie.Value
		fmt.Fprintf(w, "Hello %s", value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=daewu",nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, request)

	cookies := rec.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf( "Cookie %s : %s", cookie.Name, cookie.Value)
	}
}