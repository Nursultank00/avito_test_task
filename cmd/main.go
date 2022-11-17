package main

import (
	"log"
	"os"

	apiserver "github.com/Nursultank00/avito_test_task"
	"github.com/Nursultank00/avito_test_task/pkg/handler"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
	"github.com/Nursultank00/avito_test_task/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error occured while initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("error occured while initializing db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	srv := apiserver.NewServer(viper.GetString("port"), handler.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
