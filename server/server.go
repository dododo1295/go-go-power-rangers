package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("It started...")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		fmt.Println()
		fmt.Println()
		fmt.Println("Shutdown Executed")
		fmt.Println()
		log.Printf("Caught signal %s.\nServer shutdown complete. ", sig)
		os.Exit(0)
	}()

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":1004", nil))

}
