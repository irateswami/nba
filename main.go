package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Grab the port from the command line on startup
	// Default is 8080
	var port string
	flag.StringVar(&port, "port", "8080", "master port")
	port = ":" + port

	router := mux.NewRouter()
	//TODO
	//router.HandleFunc("/weather", METHODNAME).Methods("GET")

	// Server structure to make things a little easier
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Make a channel to kill the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// anonymous go function for listening on the server
	go func() {
		// listen and serve, if error isn't null and the server isn't closed
		// do your thing
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("server started")

	// Check the done channel if we're actually done
	<-done
	log.Print("server stopped")

	// Provide a context in which to time out
	// 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer canceling until the program exits
	defer cancel()

	// check to make sure we shutdown gracefully
	// if there is an error, this statement returns and exits here
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %+v", err)
	}

	// else, we shutdown gracefully! yay!
	log.Printf("server exited gracefully")

}
