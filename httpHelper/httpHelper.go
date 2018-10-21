package httpHelper

import "fmt"
import "net/http"

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET received")
    fmt.Println(r.Host, r.URL)
      switch r.URL.String() {
      case "/Bule":
        fmt.Println("Bule war hier")
        fmt.Fprintf(w, "<h1>Bule war hier</h1>")
      }
		break
	case http.MethodPost:
		fmt.Println("POST received")
    fmt.Println(r.Host, r.URL)
		break
	case http.MethodDelete:
		fmt.Println("DELETE received")
    fmt.Println(r.Host, r.URL)
		break
	default:
		fmt.Println("Method not allowed")
		break
	}
}
