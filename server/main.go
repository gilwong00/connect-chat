package main

import (
	"database/sql"
	"gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"
	"gilwong00/connect-chat/gen/proto/user/v1/userv1connect"
	"gilwong00/connect-chat/internal/roomservice"
	"gilwong00/connect-chat/internal/userservice"
	db "gilwong00/connect-chat/pkg/db/sqlc"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env vars", err)
		os.Exit(1)
	}
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	queries := db.New(conn)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()
	userservice, err := userservice.NewUserService(queries)
	if err != nil {
		log.Fatalln("failed to create user service")
	}
	roomservice, err := roomservice.NewRoomService(queries)
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
