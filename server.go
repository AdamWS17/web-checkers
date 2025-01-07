package main

import(
	"fmt"
	"log"
	"net/http")


func main(){	
	fmt.Println("hello world!");
	http.HandleFunc("/", callTest)
}

func callTest(w, http.ResonceWriter, r *http.Request){
	fmt.println("callTest")	
}
