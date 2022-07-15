package app

import (
	"L0/internal/repository"
	"L0/internal/server"
	"L0/internal/server/handler"
	"L0/internal/service"
	"L0/pkg/customLogger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log := customLogger.NewLogger()

	repo := repository.NewRepository(log)

	services := service.NewService(log, repo)

	handlers := handler.NewHandler(log, services)

	srv := server.NewServer(handlers.SetupRoutes(), "8000")

	if err := srv.Run(); err != nil {
		log.Panic(log.CallInfoStr(), err.Error())
	}

	exit := newExit()

	go func() {
		if err := srv.Run(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(log.CallInfoStr(), err.Error())
			}
		}
	}()

	<-exit

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(log.CallInfoStr(), err.Error())
	}

}

func newExit() chan os.Signal {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	return exit
}
