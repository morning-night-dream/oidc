package main

import (
	"context"
	"log"
	"net/http"

	"oidc/op"
	"oidc/storage"
)

func main() {
	ctx := context.Background()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	storage := storage.NewStorage(storage.NewUserStore())

	port := "9998"
	router := op.SetupServer(ctx, "http://localhost:"+port, storage)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
