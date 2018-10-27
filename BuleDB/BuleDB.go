package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "BuleDB/models"
    "BuleDB/utils"
)

//Mysql
//const (
//  host = "localhost"
//  port = "3306"
//  user = "malekith"
//  password = "laminei"
//  dbname = "test"
//)

//Postgresql
const (
  host = "localhost"
  port = "5432"
  user = "malekith"
  password = "laminei"
  dbname = "postgres"
)

func main() {

  go utils.TimeLogger()

  models.InitDB("host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
  router := mux.NewRouter()
  router.HandleFunc("/Bule", models.GetBule).Methods("GET")
  router.HandleFunc("/Bule", models.CreateBule).Methods("POST")
  router.HandleFunc("/Bule", models.DeleteBule).Methods("DELETE")
  router.HandleFunc("/Bule", models.PatchBule).Methods("PATCH")
  log.Fatal(http.ListenAndServe(":8091", router))
}
