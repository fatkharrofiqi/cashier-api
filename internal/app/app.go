package app

import (
	"cashier-api/internal/handler"
	"cashier-api/internal/repository"
	"cashier-api/internal/route"
	"cashier-api/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	// repository
	productRepo := repository.NewProductRepository()
	categoryRepo := repository.NewCategoryRepository()

	// service
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// handler
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// route
	routes := route.NewRoute(productHandler, categoryHandler)
	routes.Register()

	port := getPort()

	server := &http.Server{
		Addr: ":" + port,
	}

	return &App{
		httpServer: server,
	}
}

func (a *App) Run() {
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	go func() {
		fmt.Println("Server running on", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = a.httpServer.Shutdown(ctx)
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
