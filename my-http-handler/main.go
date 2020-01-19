package main

import "net/http"

type MyHandler struct{}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stranger..!"))
}

func main() {
	r := MyHandler{}
	http.ListenAndServe(":8080", r)
}
