package main

import(
	"fmt"
	"net/http"
)


func main(){	
	fmt.Println("Game server started...")

	http.HandleFunc("/test", callTest)

	http.HandleFunc("/", genericTest)

	http.ListenAndServe(":3000", nil)
}

func callTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("callTest")	
}

func genericTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("generic Test")
}
