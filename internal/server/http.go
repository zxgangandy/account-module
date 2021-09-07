package server

//func NewHttpServer() {
//	router := router.NewRouter()
//
//
//	server := &http.Server{
//		Addr:           ":" + viper.GetString("app.port"),
//		Handler:        router,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	startFn = func() {
//		log.Infof("http server start: %v", server.Addr)
//		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("err: %+v", err)
//		}
//	}
//
//}

import (
	"account-module/internal/router"
	"account-module/pkg/conf"
	"context"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run() {
	config := conf.GetConfig()
	router := router.Router()
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(config.Application.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithFields(logger.Fields{"err": err, "pid": syscall.Getpid()}).Fatalln("Listen server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warnln("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.WithFields(logger.Fields{"err": err}).Fatalln("Shutdown server error ...")
	}
	logger.Warnln("Server exiting")
}
