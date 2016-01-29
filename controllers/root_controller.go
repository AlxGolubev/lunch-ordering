func RootPath (w http.ResponseWriter, r *http.Request) {
  // data := UserData(db_connect)
  response := User{"Alex", "Golubev", 20}
  data, _ := json.Marshal(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}
