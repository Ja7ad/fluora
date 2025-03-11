package rest

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/healthcheck"

	_ "github.com/Ja7ad/fluora/api/rest/docs"
	"github.com/Ja7ad/fluora/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	fb     *fiber.App
	notify chan error
	ln     net.Listener
}

// NewServer
// @title Fluora API
// @version 1.0
// @description An extensible AI toolkit built in Go for generative AI applications.
// @host 127.0.0.1:8080
// @BasePath /
func NewServer(cfg *config.Rest) (*Server, error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // Avoids unnecessary log noise
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(healthcheck.New())

	corsCfg := cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}
	if len(cfg.Origins) != 0 {
		corsCfg.AllowOrigins = strings.Join(cfg.Origins, ",")
	} else {
		corsCfg.AllowOrigins = "*"
	}
	app.Use(cors.New(corsCfg))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		_, err := c.Write([]byte("Fluora server is running!"))
		if err != nil {
			return err
		}
		return nil
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	ln, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return nil, err
	}

	server := &Server{fb: app, ln: ln, notify: make(chan error)}

	return server, nil
}

func (s *Server) Start() {
	go func() {
		log.Printf("🚀 Server is running on %s", s.ln.Addr())
		s.notify <- s.fb.Listener(s.ln)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("⚠️  Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.fb.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("❌ Forced shutdown: %v", err)
	}

	log.Println("✅ Server gracefully stopped")
}

func (s *Server) Register(binder Binder, middlewares ...interface{}) {
	binder.Register(s.fb, middlewares...)
}
