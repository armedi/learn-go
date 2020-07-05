package main

import (
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // enable postgres dialect
	"google.golang.org/grpc"

	pb "github.com/armedi/learn-go/grpc/user"
	grpchandler "github.com/armedi/learn-go/handler/grpc"
	httphandler "github.com/armedi/learn-go/handler/http"
	"github.com/armedi/learn-go/user"
)

const (
	grpcPort string = ":5000"
	httpPort string = ":3000"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=armedi dbname=golearn sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}
	db.LogMode(true)
	defer db.Close()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userGrpcHandler := grpchandler.NewUserHandler(userService)
	userHTTPHandler := httphandler.NewUserHandler(userService)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// grpc api
	go func() {
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServer(s, userGrpcHandler)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
		wg.Done()
	}()

	// http api
	go func() {
		r := chi.NewRouter()
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.URLFormat)
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		r.Route("/user", func(c chi.Router) {
			r.Post("/register", userHTTPHandler.Register)
			r.Post("/login", userHTTPHandler.Login)
		})

		if err := http.ListenAndServe(httpPort, r); err != nil {
			log.Fatalf("failed to serve http: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
