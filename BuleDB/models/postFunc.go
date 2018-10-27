package models

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io"
  "io/ioutil"
  "strconv"
)

/* This function inserts a single object into the DB. Currently it assumes only a single object is transmitted and does not verify
   Something like a payload verification needs to be added
*/

func CreateBule (w http.ResponseWriter, r *http.Request) {
  var bule Bule
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
      panic(err)
  }
  if err := r.Body.Close(); err != nil {
      panic(err)
  }
  if err := json.Unmarshal(body, &bule); err != nil {
      w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
      w.WriteHeader(422) // unprocessable entity
      fmt.Fprintln(w, "Please enter a single object in the correct format")
/*      if err := json.NewEncoder(w).Encode(err); err != nil {
          panic(err)
      }*/
  } else {
    _, err = db.Exec("INSERT INTO mylittleschema.queryobjects (id, name) VALUES ($1, $2)", strconv.Itoa(bule.ID),bule.Name)
    if err != nil {
      w.Header().Set("Content-Type", "test/plain; charset=UTF-8")
      w.WriteHeader(400) // something went wrong
      fmt.Fprintln(w, err)
    }
  }
}
