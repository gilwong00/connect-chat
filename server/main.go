package main

import (
	"gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"
	"gilwong00/connect-chat/gen/proto/user/v1/userv1connect"
	"gilwong00/connect-chat/internal/roomservice"
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
	roomservice, err := roomservice.NewRoomService()
	if err != nil {
		log.Fatalln("failed to create room service")
	}
	userPath, userHandler := userv1connect.NewUserServiceHandler(userservice)
	roomPath, roomHandler := roomv1connect.NewRoomServiceHandler(roomservice)
	mux.Handle(userPath, userHandler)
	mux.Handle(roomPath, roomHandler)
	err = http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
