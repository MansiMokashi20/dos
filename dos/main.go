package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/test-server",func(rw http.ResponseWriter,r *http.Request){
		fmt.Fprintf(rw,"\nThis is test server\n")
})
http.ListenAndServe(":80",nil)}