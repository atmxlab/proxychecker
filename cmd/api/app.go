package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/atmxlab/proxychecker/cmd/app"
	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/api/proxychecker"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Services struct {
	proxychecker *proxychecker.Service
}

type App struct {
	app        *app.App
	services   Services
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp(a *app.App, checker *proxychecker.Service) *App {
	return &App{
		app: a,
		services: Services{
			proxychecker: checker,
		},
	}
}

func (a *App) Init() {
	a.app.Init()
}

func (a *App) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		logrus.Info("Starting base app...")
		if err := a.app.Start(ctx); err != nil {
			return errors.Wrap(err, "a.app.Start")
		}
		return nil
	})

	eg.Go(func() error {
		logrus.Info("Running GRPC server...")
		if err := a.runGrpcServer(ctx); err != nil {
			return errors.Wrap(err, "a.runGrpcServer")
		}
		return nil
	})

	eg.Go(func() error {
		logrus.Info("Running HTTP server...")
		if err := a.runHTTPServer(ctx); err != nil {
			return errors.Wrap(err, "a.runHTTPServer")
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "eg.Wait()")
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.grpcServer.GracefulStop()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "a.httpServer.Shutdown")
	}

	return nil
}

func (a *App) runGrpcServer(ctx context.Context) error {
	cfg := a.app.Container().Config()
	logrus.Infof("Listening GRPC port: [%d]", cfg.API.GRPC.Port)

	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%d", cfg.API.GRPC.Port),
	)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}

	a.grpcServer = grpc.NewServer()
	desc.RegisterProxycheckerServer(a.grpcServer, a.services.proxychecker)
	reflection.Register(a.grpcServer)

	// TODO graceful stop
	if err = a.grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "grpcServer.Serve")
	}

	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	cfg := a.app.Container().Config()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := desc.RegisterProxycheckerHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf(":%d", cfg.API.GRPC.Port),
		opts,
	); err != nil {
		return errors.Wrap(err, "RegisterProxycheckerHandlerFromEndpoint")
	}

	httpMux := http.NewServeMux()

	httpMux.Handle("/", mux)

	logrus.Infof(
		"Setup swagger JSON handler: url: [%s], path: [%s]",
		cfg.API.HTTP.Swagger.JSON.URL,
		cfg.API.HTTP.Swagger.JSON.Path,
	)

	httpMux.HandleFunc(cfg.API.HTTP.Swagger.JSON.URL, func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Get swagger JSON endpoint")
		http.ServeFile(w, r, cfg.API.HTTP.Swagger.JSON.Path)
	})

	logrus.Infof(
		"Setup swagger UI handler: url: [%s], sourcePath: [%s]",
		cfg.API.HTTP.Swagger.UI.URL,
		cfg.API.HTTP.Swagger.UI.SourcePath,
	)

	httpMux.Handle(
		cfg.API.HTTP.Swagger.UI.URL,
		http.StripPrefix(
			cfg.API.HTTP.Swagger.UI.URL,
			http.FileServer(http.Dir(cfg.API.HTTP.Swagger.UI.SourcePath)),
		),
	)

	logrus.Infof("Setup CORS middleware")
	handler := a.corsMiddleware(httpMux)

	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.API.HTTP.Port),
		Handler: handler,
	}
	if err := a.httpServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "net.ListenAndServe")
	}

	return nil
}

func (a *App) corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любого источника
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Обрабатываем preflight запросы
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
