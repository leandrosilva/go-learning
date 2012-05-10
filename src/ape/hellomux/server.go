package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var homeHtmlTemplate = template.Must(template.New("").Parse(homeHtmlTemplateStr))
var helloHtmlTemplate = template.Must(template.New("n").Parse(helloHtmlTemplateStr))

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/hello/{name}", helloHandler)

	fmt.Println("Hellomux server listening at 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("[homeHandler] Request:", req)

	homeHtmlTemplate.Execute(res, nil)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("[helloHandler] Request:", req)

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
