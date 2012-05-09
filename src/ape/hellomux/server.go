package main

import (
  "fmt"
  "log"
  "net/http"
  "text/template"
  "code.google.com/p/gorilla/mux"
)

var homeHtmlTemplate = template.Must(template.New("").Parse(homeHtmlTemplateStr))
var helloHtmlTemplate = template.Must(template.New("n").Parse(helloHtmlTemplateStr))

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", HomeHandler)
  router.HandleFunc("/hello/{name}", HelloHandler)
  
  fmt.Println("Hellomux server listening at 8080")
  
  err := http.ListenAndServe(":8080", router)
  if err != nil {
      log.Fatal("ListenAndServe:", err)
  }
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
  log.Println("[HomeHandler] Request:", req)

  homeHtmlTemplate.Execute(res, nil)
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {
  log.Println("[HelloHandler] Request:", req)
  
  params := mux.Vars(req)
  name := params["name"]

  helloHtmlTemplate.Execute(res, name)
}

const homeHtmlTemplateStr = `
<html>
<head>
  <title>go-learning :: ape :: hellomux - home</title>
</head>
<body>
  This is my third program in Go and it is the first one using Gorilla Mux. So far, everything is awesome.
</body>
</html>
`
const helloHtmlTemplateStr = `
<html>
<head>
  <title>go-learning :: ape :: hellomux - say hello</title>
</head>
<body>
  Hello {{html .}}!!!
</body>
</html>
`
