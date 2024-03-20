package gapi

import (
	"context"
	"time"

	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/dpurbosakti/booknest-grpc/internal/pb"
	"github.com/dpurbosakti/booknest-grpc/internal/util"
	"github.com/dpurbosakti/booknest-grpc/internal/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Name:           req.GetName(),
			HashedPassword: hashedPassword,
			Phone:          req.GetPhone(),
			Email:          req.GetEmail(),
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Name,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)

		},
	}

	// log.Info().Msg(">> creating user ...")
	// time.Sleep(10 * time.Second)
	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}

	// log.Info().Msg(">> done creating user.")
	return rsp, nil
}
