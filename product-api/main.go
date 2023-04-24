package main

import (
	"context"


	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"

	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)
var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind Address")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server")
var basePath = env.String("BASE_PATH", false, "/tmp/images", "Base path to save images")

func main() {
	env.Parse()
	l := hclog.New(
		&hclog.LoggerOptions{
			Name: "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)
	sl:= l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})
	//stor, err :=

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
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	corshandle := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	ph := serveMux.Get(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}")
	s := &http.Server{
		Addr:         ":9090",
		Handler:      corshandle(serveMux),
		ErrorLog: sl,
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
