package main

import (
	"fmt"
	goredis "github.com/go-redis/redis"
	"gitlab.com/M4xwell/network-service/host"
	"gitlab.com/M4xwell/network-service/pkg/config"
	"gitlab.com/M4xwell/network-service/pkg/logger"
	"gitlab.com/M4xwell/network-service/redis"
	"gitlab.com/M4xwell/network-service/server"
	"time"
)

func main() {
	cfg := config.NewConfig()
	log := *logger.NewLogger()

	log.Info("Start server")

	client := initRedis(cfg)

	// redis datasource
	repository := redis.Repository{
		Logger: log,
		Client: *client,
	}

	// domain logic
	service := host.Service{
		Repository: &repository,
		Logger:     log,
	}

	// gets hosts from nmap
	job := host.Job{
		Repository: &repository,
		Logger:     log,
		Target:     cfg.NmapTarget,
	}

	go scheduleJob(job)

	srv := server.NewServer(log, service)

	srv.Serve()
}

func initRedis(cfg config.Config) *goredis.Client {
	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	return client
}

func scheduleJob(job host.Job) {
	for true {
		_ = job.FindHosts()

		time.Sleep(5 * time.Minute)
	}
}
