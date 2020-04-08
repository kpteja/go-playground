package main

import "net/http"

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}

func main() {
	http.Handle("/", homeHandler())
}
