package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareOne [BEFORE]")
		next.ServeHTTP(w, r)
		log.Println("middlewareOne [AFTER]")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middlewareTwo [BEFORE]")
		next.ServeHTTP(w, r)
		log.Println("middlewareTwo [AFTER")
	})
}

func finalHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("finalHandler")
		w.Write([]byte("Final Handler..!"))
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome..!"))
	})

	http.Handle("/handle", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handle..!"))
	}))

	http.Handle("/middleware", middlewareOne(middlewareTwo(finalHandler())))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
