package main

import (
	"gilwong00/connect-chat/gen/proto/user/v1/userv1connect"
	"gilwong00/connect-chat/internal/userservice"
	"log"

	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	userservice, err := userservice.NewUserService()
	if err != nil {
		log.Fatalln("failed to create user service")
	}
	userPath, userHandler := userv1connect.NewUserServiceHandler(userservice)
	mux.Handle(userPath, userHandler)
	err = http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
