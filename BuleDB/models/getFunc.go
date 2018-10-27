package models

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
)


/*
This function will retrieve records from the database table.
The current implementation is a bit crude and allows only retrieving one or all data sets,
this should be generalized sometime to retrieve any number of entries

*/
func GetBule (w http.ResponseWriter, r *http.Request) {
  var (
    id int
    name string
  )

  params := r.URL.Query()["id"]
  // check how many query parameters of type "ID" were included
  // if it is none, return all objects from the DB table

  if len(params) == 0 {
    var resultSet []Bule
    rows, err := db.Query("SELECT * FROM mylittleschema.queryobjects ORDER BY id")
    if err != nil {
  	   log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
     err := rows.Scan(&id, &name)
     if err != nil {
      log.Fatal(err)
     } else {
       resultSet = append(resultSet, Bule{ID: id, Name: name})
     }
    }
    json.NewEncoder(w).Encode(resultSet)

    // else, if exactly one parameter has been given, fetch exactly this entry
  } else if len(params) == 1 {
    _, err := strconv.Atoi(params[0])
    if err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(400) // Bad Request, the ID was not a number
    } else {
      // This should return either one or no rows, as "ID" is the primary key on the table.
      rows, err := db.Query("select * from mylittleschema.queryobjects where id = $1", params[0])
      if err != nil {
    	   log.Fatal(err)
      }
      defer rows.Close()

      if rows.Next() == true {
        err := rows.Scan(&id, &name)
        if err != nil {
          log.Fatal(err)
        } else {
        json.NewEncoder(w).Encode(Bule{ID: id, Name: name})
        }
      } else {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(404) // No record found with that ID
      }
    }
  }
}

type Bule struct {
  ID int
  Name string
}

var bules []Bule
