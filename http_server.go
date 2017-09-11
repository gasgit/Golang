package main

import(
	"fmt"
	"io"
	"net/http"
)

func sayHello(res http.ResponseWriter, req *http.Request){
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res , "<html><head><title>Hello</title></head><body > Hello GO </body></html>",	
	)
}

func main(){
	http.HandleFunc("/", sayHello)
	fmt.Println("Running on http://127.0.0.1:9000")
	http.ListenAndServe(":9000",nil)
	

}