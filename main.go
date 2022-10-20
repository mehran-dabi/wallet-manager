package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet-manager/config"
	"wallet-manager/domain/wallet/controller"
	"wallet-manager/domain/wallet/repository"
	"wallet-manager/domain/wallet/service"
	"wallet-manager/infrastructure/database"
)

func main() {
	//init config
	conf := config.Init()

	// init db
	store, err := database.NewDatabase(
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.Driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	// check database health
	if err := store.Ping(); err != nil {
		log.Fatalf("failed to get database ping: %s", err)
	}

	// create tables if they don't exist
	if err := store.Migrate("up"); err != nil {
		log.Fatalf("failed to migrate the schemas: %s", err)
	}

	walletsRepo := repository.NewWalletRepository(store.DB())
	walletsService := service.NewWalletService(walletsRepo)
	walletsController := controller.NewWalletController(walletsService, store)

	server := walletsController.Run(conf.Service.Port)

	waitForOsSignal()
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func waitForOsSignal() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-osSignal
}
