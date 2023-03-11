package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Olá!"))
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Request processada com sucesso")     // log do server
		w.Write([]byte("Request processada com sucesso")) // imprime no browser
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
