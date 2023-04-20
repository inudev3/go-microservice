package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product", log.LstdFlags)

	products := handlers.NewProducts(l)

	serveMux := mux.NewRouter()
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	addRouter := serveMux.Methods(http.MethodPost).Subrouter()
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	getRouter.HandleFunc("/", products.GetProducts)
	addRouter.HandleFunc("/product", products.AddProduct)
	addRouter.Use(products.MiddlewareValidation)
	putRouter.HandleFunc("/product/{id:[0-9]+}", products.UpdateProduct)
	putRouter.Use(products.MiddlewareValidation)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
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
	l.Println("Received terminate, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(30))

	s.Shutdown(ctx)
}
