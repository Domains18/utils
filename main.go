package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 nor found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request succesful")
	name:= r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}


func main(){
	var port int = 3000
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port", port)
	if err := http.ListenAndServe(":" + strconv.Itoa(port), nil); err!= nil{
		log.Fatal(err)
	}
}
