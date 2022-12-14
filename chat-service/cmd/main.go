package main

import (
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc"
	"github.com/widcraft/chat-service/internal/adapter/repository"
	"github.com/widcraft/chat-service/internal/adapter/repository/redis"
	"github.com/widcraft/chat-service/internal/adapter/rest"
	chatapp "github.com/widcraft/chat-service/internal/application/chat"
)

var logger = log.New()
var redisDb *redis.Redis

var (
	restServer *rest.Rest
	grpcServer *grpc.Grpc
)

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

// DB
func init() {
	var err error
	redisDb, err = redis.New(logger, net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)

	if err != nil {
		shutdown(redisDb)
		logger.Fatalf("redis connection failure: %s", err)
	}
}

// servers
func init() {
	chatRepo := repository.NewChatRepository(logger, redisDb)
	chatApp := chatapp.New(logger, chatRepo)

	restServer = rest.New(logger, chatApp)
	grpcServer = grpc.New(logger, chatApp)
}

func main() {
	defer gracefulShutdown()
	go restServer.Run(os.Getenv("REST_PORT"))
	go grpcServer.Run(os.Getenv("GRPC_PORT"))
}

func gracefulShutdown() {
	defer shutdown(restServer, grpcServer, redisDb)

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(targets ...interface{ Close() error }) {
	for _, target := range targets {
		if !reflect.ValueOf(target).IsNil() {
			if err := target.Close(); err != nil {
				logger.Errorf("%s closing failed: %s", reflect.TypeOf(target), err)
			}
		}
	}
}
