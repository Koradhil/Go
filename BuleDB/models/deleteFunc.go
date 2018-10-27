package models

import (
  "net/http"
  "fmt"
)

func DeleteBule (w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query()["id"]
  if len(id) == 1 {
    //One ID has been set, find the right element and return it
    _, err := db.Exec("DELETE FROM mylittleschema.queryobjects where id = $1", id[0])
    if err != nil {
      w.Header().Set("Content-Type", "test/plain; charset=UTF-8")
      w.WriteHeader(400) // something went wrong
      fmt.Fprintln(w, err)
    }
  } else {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(400) // Bad Request
    fmt.Fprintln(w, "Please enter a single ID for the object as parameter")
  }
}
