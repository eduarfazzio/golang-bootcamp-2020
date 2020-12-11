package main

import (
	"net/http"
)

type variables struct {
	Name string `json:"name"`
}

func hello(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//REST
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
