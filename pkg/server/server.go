package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/takutakahashi/any-exporter/pkg/config"
	"github.com/takutakahashi/any-exporter/pkg/metrics"
)

type Server struct {
	config  config.Config
	metrics metrics.MetricsStore
}

func New(configPath string) (Server, error) {
	c, err := config.Load(configPath)
	if err != nil {
		return Server{}, err
	}
	return Server{
		config:  c,
		metrics: metrics.NewMetricsStore(),
	}, nil
}

func (s Server) StartServer() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healty")

	})
	e.GET("/metrics", func(c echo.Context) error {
		return c.String(http.StatusOK, s.metrics.String())
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Start(s.Address())
}

func (s Server) StartWorker() {
	for _, mc := range s.config.Metrics {
		m := metrics.NewMetrics(mc)
		s.metrics.Add(m)
	}
	logrus.Info(s.metrics.Length())
	for {
		err := s.metrics.ExecuteAll()
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(s.metrics)
		time.Sleep(30 * time.Second)
		s.Sleep()
	}
}

func (s Server) Sleep() {
	duration, err := time.ParseDuration(s.config.Resolution)
	if err != nil {
		duration = 30 * time.Second
	}
	time.Sleep(duration)
}

func (s Server) Address() string {
	var host, port string
	if s.config.Host == "" {
		host = "localhost"
	} else {
		host = s.config.Host
	}
	if s.config.Port == "" {
		port = "9400"
	} else {

		port = s.config.Port
	}

	return host + ":" + port
}
