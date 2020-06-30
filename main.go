package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // enable postgres dialect

	"github.com/armedi/learn-go/handler"
	"github.com/armedi/learn-go/user"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=armedi dbname=golearn sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Route("/user", func(c chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})

	http.ListenAndServe(":3000", r)
}
