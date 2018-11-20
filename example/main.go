package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greetings string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greetings)))
}
func main() {
	http.Handle("/", &myHandler{"Hello"})
	http.ListenAndServe(":8000", nil)
}
