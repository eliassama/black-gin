package api

import (
	"context"
	"fmt"
	"github.com/eliassama/black-gin/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Opts struct {
	Route  *gin.Engine
	Listen string
	Logger *zap.Logger
}

type Service struct {
	svr    *http.Server
	logger *zap.Logger
}

func New(opts *Opts) *Service {
	if opts.Route == nil {
		panic("Route Not Found")
	}

	if opts.Logger == nil {
		panic("Logger Not Found")
	}

	if opts.Listen == "" {
		opts.Listen = "0.0.0.0:9876"
	}

	middleware.Use(opts.Route, opts.Logger)

	service := &Service{
		svr: &http.Server{
			Addr:    opts.Listen,
			Handler: opts.Route,
		},
		logger: opts.Logger,
	}

	return service
}

func (s Service) Start() error {
	if err := checkPortAvailable(s.svr.Addr); err != nil {
		s.logger.Error(fmt.Sprintf("port %s is already in use: %s", s.svr.Addr, err))
		return err
	}

	go func() {
		if err := s.svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(fmt.Sprintf("server listen err:%s", err))
		}
	}()

	addrSlice := strings.Split(s.svr.Addr, ":")
	port := addrSlice[len(addrSlice)-1]

	localIP, err := getLocalIP()
	if err != nil {
		s.logger.Error(fmt.Sprintf("error getting local IP: %s", err))
	} else {
		s.logger.Info(fmt.Sprintf("Service Is Running In:\n http://%s:%s \n http://127.0.0.1:%s\n", localIP, port, port))
	}

	s.logger.Info("Waiting for request...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTSTP, syscall.SIGTERM)

	<-quit

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)

	defer channel()
	if err := s.svr.Shutdown(ctx); err != nil {
		s.logger.Panic("server shutdown error")
	}
	s.logger.Info("server exiting...")

	return nil
}
