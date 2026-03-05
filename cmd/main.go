package main

import (
	"context"
	todo "my_api"
	"my_api/pkg/handler"
	"my_api/pkg/repository"
	"my_api/pkg/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatal("error loading .env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		SSLMode:  viper.GetString("db.sslmode"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("fail to initialize db: %s ", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
			log.Fatalf("error ocured while running http server: %s", err.Error())
		}
	}()
	log.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Print("TodoApp Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error shutting down http server: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("error closing db: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
