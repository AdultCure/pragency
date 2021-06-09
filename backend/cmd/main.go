package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	backend "pragency"
	"pragency/pkg/handler"
	"pragency/pkg/repository"
	"pragency/pkg/service"
	"syscall"
)

// Функция запуска всего приложения

func main() {

	// Загрузка конфига

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializationg configs:  %s", err.Error())
	}

	// Загруска env

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Подключение к дб

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	// Ошибка подключения к бд

	if err != nil {
		logrus.Fatalf("Failed to initialize database: %s", err.Error())
	}

	// Указатели

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Запуск сервера в Горутине

	srv := new(backend.Server)
	go func () {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	// Уведомление запуска сервера

	logrus.Print("PR agency app Started")

	// Выход из Горутины и остановки сервера

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	// Уведомление остановки сервера

	logrus.Print("PR agency Shutting Down")

	// Ошибка остановки сервера

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	// Ошибка закрытия бд

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

// Функция инициации конфига в viper

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
