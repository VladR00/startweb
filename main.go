package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := InitiateTables(); err != nil {
		log.Fatalf("%v", err)
	}

	http.HandleFunc("/registration", HandleRegistrationFunc)
	http.HandleFunc("/login", HandleLoginFunc)

	fmt.Println("Server start at 8008")
	log.Fatal(http.ListenAndServe(":8008", nil))
	select {}
}
