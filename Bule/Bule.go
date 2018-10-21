package main

import (
    "encoding/json"
    "log"
    "net/http"
    "io"
    "io/ioutil"
    "github.com/gorilla/mux"
)

func main() {

    bules = append(bules, Bule{ID: 1, Name: "TheRealBule"} )
    router := mux.NewRouter()
    router.HandleFunc("/Bule", GetBule).Methods("GET")
    router.HandleFunc("/Bule", CreateBule).Methods("POST")
    log.Fatal(http.ListenAndServe(":8091", router))
}

func GetBule (w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(bules)
}

func CreateBule (w http.ResponseWriter, r *http.Request) {
  log.Println(r.Body)
  var bule Bule
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
      panic(err)
  }
  if err := r.Body.Close(); err != nil {
      panic(err)
  }
  if err := json.Unmarshal(body, &bule); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(422) // unprocessable entity
      if err := json.NewEncoder(w).Encode(err); err != nil {
          panic(err)
      }
  }
  bules = append(bules, bule)
}

type Bule struct {
  ID int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
}

var bules []Bule
