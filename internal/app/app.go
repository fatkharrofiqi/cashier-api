package app

import (
	"cashier-api/internal/config/config"
	"cashier-api/internal/config/db"
	"cashier-api/internal/handler"
	"cashier-api/internal/repository"
	"cashier-api/internal/route"
	"cashier-api/internal/service"
	"cashier-api/internal/uow"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewApp() *App {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConn, err := db.NewDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// repository
	productRepo := repository.NewProductRepository(dbConn)
	categoryRepo := repository.NewCategoryRepository(dbConn)
	transactionRepo := repository.NewTransactionRepository(dbConn)
	transactionDetailRepo := repository.NewTransactionDetailRepository(dbConn)
	reportRepo := repository.NewReportRepository(dbConn)

	// service
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	reportService := service.NewReportService(reportRepo)

	// uow
	unitOfWork := uow.NewUnitOfWork(dbConn)

	// transaction service with uow
	transactionService := service.NewTransactionService(productRepo, transactionRepo, transactionDetailRepo, unitOfWork)

	// handler
	productHandler := handler.NewProductHandler(productService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	reportHandler := handler.NewReportHandler(reportService)

	// route
	routes := route.NewRoute(productHandler, categoryHandler, transactionHandler, reportHandler)
	routes.Register()

	server := &http.Server{
		Addr: ":" + cfg.Port,
	}

	return &App{
		httpServer: server,
		db:         dbConn,
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
