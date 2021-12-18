package belajar_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w,"simple.gohtml", "Hello template HTML")
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	rec := httptest.NewRecorder()

	SimpleHtmlFile(rec, request)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}