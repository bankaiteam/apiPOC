package main

import (
  "net/http"
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/error", func() (int, string) {
  return 400, "i'm a teapot" // HTTP 418 : "i'm a teapot"
  })

  m.Get("/params", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
  res.WriteHeader(200) // HTTP 200
  })
  m.Run()
}