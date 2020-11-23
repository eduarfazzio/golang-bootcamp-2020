package main

import "net/http"

func hello (w http.ResponseWriter, r *http.Request){

}

func main (){
	http.HandleFunc ("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
	panic(err)
	}
}