package main

import "fmt"
import "net/http"
import "httpHelper"

func main(){
  fmt.Println("Starting up")

  fmt.Println("Attempting to start HTTP Server.")

	http.HandleFunc("/", httpHelper.HandleRequest)

  var err = http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Println("Server failed starting. Error: %s", err)
	}
}
