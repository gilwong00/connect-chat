package main

import (
	"context"
	"database/sql"
	"errors"
	"gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"
	"gilwong00/connect-chat/gen/proto/user/v1/userv1connect"
	"gilwong00/connect-chat/internal/roomservice"
	"gilwong00/connect-chat/internal/userservice"
	db "gilwong00/connect-chat/pkg/db/sqlc"
	"gilwong00/connect-chat/pkg/ws"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/lib/pq"
)

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			// In prod be exclusive
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env vars", err)
		os.Exit(1)
	}
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()
	queries := db.New(conn)
	hub := ws.NewHub()
	hubHandler := ws.NewHubHandler(hub)
	userservice, err := userservice.NewUserService(queries)
	if err != nil {
		log.Fatalln("failed to create user service")
	}
	roomservice, err := roomservice.NewRoomService(queries, hubHandler)
	if err != nil {
		log.Fatalln("failed to create room service")
	}
	userPath, userHandler := userv1connect.NewUserServiceHandler(userservice)
	roomPath, roomHandler := roomv1connect.NewRoomServiceHandler(roomservice)
	mux.Handle(userPath, userHandler)
	mux.Handle(roomPath, roomHandler)
	// websocket only path
	mux.Handle("/room/join", hubHandler)
	// calling the Run method in our Hub on a separate go routine
	go hub.Run()
	// new server
	srv := &http.Server{
		Addr: "localhost:8080",
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err)
	}
	// old setup
	// err = http.ListenAndServe(
	// 	"localhost:8080",
	// 	h2c.NewHandler(
	// 		newCORS().Handler(mux),
	// 		&http2.Server{},
	// 	),
	// )
	// log.Fatalf("listen failed: %v", err)
}
