package main

import (
  // "fmt"
  "net/http"
  "log"
  "encoding/json"
  "database/sql"
  "../models"
  "../controllers"

  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)



func db_connect() {
  db, err := sql.Open("postgres", "user=alxgol dbname=lunch_ordering_deveelopment")
  if err != nil {
    log.Fatal(err)
    return err
  } else {
    return db
  }
}

func UserData(db) {
  rows, err = db.QueryRow(`SELECT * FROM users`)
  if err != nil {
    log.Fatal(err)
    return err
  } else {
    return rows
  }
}

func main () {
  db := db_connect
  router := mux.NewRouter()
  router.HandleFunc("/", RootPath).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))
}
