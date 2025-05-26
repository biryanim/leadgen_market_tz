package app

import (
	"context"
	"github.com/biryanim/leadgenmarket_tz/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	go func() {
		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run http server: %v", err)
		}

		_ = a.serviceProvider.RedisClient().Close()
	}()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("local.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	router := gin.Default()

	public := router.Group("/")
	{
		public.POST("/encode", a.serviceProvider.ApiImpl(ctx).EncodeText)
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: router,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
