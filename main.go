package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/endanga/product-service/database"
	"github.com/endanga/product-service/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "admin123"
	dbname := "Product"

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database.DBCon, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("error : ", err)
	}

	defer database.DBCon.Close()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)

	getMaxRouter := sm.Methods("GET").Subrouter()
	getMaxRouter.HandleFunc("/products/max", ph.GetMaxRowCount)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
