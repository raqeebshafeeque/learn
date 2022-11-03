package main

import (
	"fmt"
	"learntest/maths"
	"learntest/server"
	"log"
	"net/http"
)

func main() {
	sum := maths.Add(5, 5)
	fmt.Printf("5 + 5 = %d\n", sum)

	http.HandleFunc("/double", server.DoubleHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The server is listening on port 8080")
}
