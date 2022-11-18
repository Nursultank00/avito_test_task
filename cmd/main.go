package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	apiserver "github.com/Nursultank00/avito_test_task"
	"github.com/Nursultank00/avito_test_task/pkg/handler"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
	"github.com/Nursultank00/avito_test_task/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error occured while initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatalf("error occured while initializing db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	srv := apiserver.NewServer(viper.GetString("port"), handler.InitRoutes())

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
