package main

import (
	"context"
	"fmt"
	"subscription_tracker/pkg/application"
	"subscription_tracker/pkg/server"
)

func main() {
	ctx := context.Background()
	app, err := application.New(ctx)
	if err != nil {
		fmt.Println("Error creating application")
		return
	}

	srv, err := server.New(app)
	if err != nil {
		fmt.Println("Error creating server")
		return
	}

	//Start the server
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server")
	}

}
