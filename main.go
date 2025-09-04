package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thomzes/fem-project/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("we are running our app")

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Panic(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is health check func!")
}
