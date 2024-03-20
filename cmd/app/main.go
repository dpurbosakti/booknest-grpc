package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dpurbosakti/booknest-grpc/internal/api"
	"github.com/dpurbosakti/booknest-grpc/internal/config"
	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/dpurbosakti/booknest-grpc/internal/gapi"
	"github.com/dpurbosakti/booknest-grpc/internal/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.ENVIRONMENT == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(connPool)

	waitGroup, ctx := errgroup.WithContext(ctx)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	api.RunEchoServer(ctx, waitGroup, config, store)
	gapi.RunGatewayServer(ctx, waitGroup, config, store, taskDistributor)
	gapi.RunGRPCServer(ctx, waitGroup, config, store, taskDistributor)
	worker.RunTaskProcessor(ctx, waitGroup, config, redisOpt, store)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}
