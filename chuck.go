package main

import (
  "io"
  "net/http"
  "encoding/json"
  "os"
  "fmt"
)

type Joke struct {
  Category  string
  Icon      string
  Id        string
  Url       string
  Value     string
}

func getJson(url string, target interface{}) error {
  r, err := http.Get(url)
  if err != nil {
    return err
  }
  defer r.Body.Close()
  return json.NewDecoder(r.Body).Decode(target)
}

func getJoke(w http.ResponseWriter, r *http.Request) {
  joke := Joke{}
  getJson("https://api.chucknorris.io/jokes/random", &joke)
  io.WriteString(w, joke.Value)
}

func main() {
  http.HandleFunc("/", getJoke)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }
  fmt.Println("Listening on port: ", port)
  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
    panic(err)
  }
}
