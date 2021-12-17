package belajar_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world")
	}
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Selamat Datang")
	})
	mux.HandleFunc("/daewu/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Selamat Datang Daewu")
	})
	mux.HandleFunc("/daewu/gus/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Selamat Datang Daewu Gus")
	})
	mux.HandleFunc("/daewu/gus/bintara", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Halo Selamat Datang Daewu Gus Bintara Putra")
	})
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}