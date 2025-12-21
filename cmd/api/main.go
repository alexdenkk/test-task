package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	"os"
	"time"

	"net/http"
	"github.com/gorilla/mux"

	"alexdenkk/test-task/internal/handler"
	"alexdenkk/test-task/internal/model"
	"alexdenkk/test-task/internal/repository"
)

func main() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("database connected: ", os.Getenv("DB_NAME"))

	db.AutoMigrate(&model.Number{})
	log.Println("migration successful")

	repo := repository.NewRepository(db)
	log.Println("repository instance created")

	h := handler.NewHandler(repo)
	log.Println("http handler instance created")

	router := mux.NewRouter()
	log.Println("router instance created")

	router.HandleFunc("/", h.AddNumber).Methods("POST")
	log.Println("routes registered")

	srv := &http.Server{
		Handler: router,
		Addr:    os.Getenv("HOST"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("running on host: ", os.Getenv("HOST"))
	log.Fatalln(srv.ListenAndServe())
}
