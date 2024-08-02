package main

import {
	"fmt"
	"net/http"
}

func main(){
	http.HandleFUnc("/", HelloServer)
	http.listenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request){
	fmt.fPrint(w, "Hello World")
}