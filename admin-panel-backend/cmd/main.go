package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hudayberdyyev/admin-panel-backend/pkg/handler"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/repository"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/service"
	"github.com/Hudayberdyyev/admin-panel-backend/server"
	"github.com/Hudayberdyyev/admin-panel-backend/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	dbPool, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     uint16(viper.GetInt("db.port")),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	defer dbPool.Close()

	if err != nil {
		logrus.Fatalf("error initializing database: %s", err.Error())
	}

	storage.ImageStorage, err = storage.NewNewsStorage(storage.MinioConfig{
		Endpoint:       viper.GetString("storage.host") + ":" + viper.GetString("storage.port"),
		AccessKeyId:    viper.GetString("storage.username"),
		SecretAccesKey: viper.GetString("storage.password"),
		UseSSL:         viper.GetBool("storage.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error initializing minio storage: %s", err.Error())
	}

	repos := repository.NewRepository(dbPool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.ip"), viper.GetString("server.port"), viper.GetString("server.protocol"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Server started ...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Print("Server Shutting Down")

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
