package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/wagaru/task/config"
	"github.com/wagaru/task/internal/delivery"
	"github.com/wagaru/task/internal/repository"
	"github.com/wagaru/task/internal/service"
)

var (
	configFile string
	conf       *config.ServerConfig
)

func init() {
	var err error
	if err = setupFlag(); err != nil {
		log.Fatalf("init.setupFlag err:%v", err)
	}
	if err = setupConfig(); err != nil {
		log.Fatalf("init.setupConfig err:%v", err)
	}
}

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	delivery := delivery.NewDelivery(svc, conf)
	go func() {
		if err := delivery.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server run error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("start shut down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := delivery.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("server exiting...")

}

func setupFlag() error {
	flag.StringVar(&configFile, "c", "../config", "設定檔路徑")
	flag.Parse()
	return nil
}

func setupConfig() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("Server", &conf); err != nil {
		return err
	}
	if port := viper.GetString("port"); port != "" {
		conf.Port = port
	}
	if mode := viper.GetString("mode"); mode != "" {
		conf.RunMode = mode
	}
	return nil
}
