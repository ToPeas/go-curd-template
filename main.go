package main

import (
	"context"
	"fmt"
	_ "go.uber.org/automaxprocs"
	"log"
	"github/ToPeas/go-curd-templatemysql"
	"github/ToPeas/go-curd-templatepkg/setting"
	"github/ToPeas/go-curd-templatepkg/validator"
	"github/ToPeas/go-curd-templaterouters"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	setting.Setup()
	mysql.Setup()
}

func main() {
	router := routers.InitRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Config.App.Port),
		Handler: router,
	}

	validator.InitValidator()
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Server running at http://127.0.0.1:%d", setting.Config.App.Port)
		// serve connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
