package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Ja7ad/fluora/version"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ja7ad/fluora/api/rest/controller"
	"github.com/Ja7ad/fluora/config"
	"github.com/Ja7ad/fluora/internal/service"
	"github.com/Ja7ad/fluora/pkg/logger"

	"github.com/Ja7ad/fluora/api/rest"
	"github.com/Ja7ad/fluora/internal/adapter/ai"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfgPath := flag.String("config", "", "Path to config file")
	ver := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *ver {
		fmt.Println("Fluora " + version.Version.String())
		os.Exit(0)
	}

	cfg, err := config.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.BasicCheck(); err != nil {
		log.Fatal(err)
	}

	logger.InitGlobalLogger(cfg.Logger)

	ai, err := ai.NewAIFromConfig(ctx, cfg.AI)
	if err != nil {
		logger.Fatal("failed to init ai", "err", err)
	}

	textcraftSvc := service.NewTextCraft(ai)

	textcraftHandler := controller.NewTextCraftController(textcraftSvc)

	restServer, err := rest.NewServer(cfg.Rest)
	if err != nil {
		log.Fatalf("error creating server: %v", err)
	}

	handlerMiddlewares := make([]any, 0)

	restServer.Register(textcraftHandler, handlerMiddlewares...)
	restServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		restServer.GracefulShutdown()
		logger.Warn("app/run: signal received", "signal", s.String())
	case err := <-restServer.Notify():
		logger.Error("rest server got error", "err", err)
	}
}
