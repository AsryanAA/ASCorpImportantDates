package main

import (
	docs "ASCorpImportantDates/api"
	"ASCorpImportantDates/internal/config"
	"ASCorpImportantDates/internal/http_server/handlers"
	"ASCorpImportantDates/internal/lib/logger"
	"ASCorpImportantDates/internal/lib/mw"
	"ASCorpImportantDates/internal/storage/sqlite"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config (cleanenv)
	cfg := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(cfg.Env)
	log.Info("starting ASCorpImportantDates", slog.String("env", cfg.Env))
	log.Debug("debug messages")

	// TODO: init storage: SQLite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}

	// TODO: init router: gin
	router := gin.Default()

	router.POST("auth", handlers.SignIn(storage))
	router.POST("reg", handlers.CreateUser(storage))

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			users := v1.Group("users")
			{
				users.GET("read", handlers.ReadUser("alphocap", storage))
				users.GET("all", mw.CheckHeader("iva"), handlers.ReadUsers(storage))
			}
		}
	}

	docs.SwaggerInfo.Title = "Автодокументация к приложению ASCorpImportantDates"
	docs.SwaggerInfo.BasePath = "api/v1"

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/views/css", "templates/css")

	// TODO: run server
	err = router.Run(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port))
	if err != nil {
		log.Error("can not start server", logger.Err(err))
		os.Exit(1)
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
