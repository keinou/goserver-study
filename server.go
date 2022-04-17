package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	fmt.Printf("Starting server at port 80\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

/* Usado para forcar erro sonarcloud */
func forError() {
	var b = 1

	if b == 0 {
		log.Fatal("err")
	} else {
		log.Fatal("err")
	}
}

/* Usado para forcar problemas de segurança */
// func forSecurity() {
// 	var (
// 		ip   = "192.168.12.42"
// 		port = 3333
// 	)

// 	log.Fatal(ip)
// 	log.Fatal(port)
// }

// /* Usado para forcar melhoria de codigo */
// func forCodeSmell1() {
// 	log.Fatal("Alguma coisa duplicada aqui!!!")
// 	log.Fatal("Alguma coisa duplicada aqui!!!")
// 	log.Fatal("Alguma coisa duplicada aqui!!!")
// }

// func forCodeSmell2() {
// 	//FIXME
// }
