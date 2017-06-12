package main

// import reqiured packages
import(
  "fmt"
  "net/http"
)

// create  handler function with ResponseWriter and Request

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello %s!\n", r.URL.Path[1:])

}

// run and test through browser "http://localhost:8080/<your-name-here>"
// or
// curl -X GET http://localhost:8080/<your-name-here>

func main(){

  fmt.Println("server running on http://localhost:8080/<add-your-name-here>")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)

}
