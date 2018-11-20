package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greetings string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("hello world", )))
	})
	http.ListenAndServe(":8000", nil)
}
