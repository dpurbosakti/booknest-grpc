package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dpurbosakti/booknest-grpc/internal/util"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSaveToExcel = "task:save_to_excel"

type PayloadSaveToExcel struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSaveToExcel(
	ctx context.Context,
	payload *PayloadSaveToExcel,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSaveToExcel, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSaveToExcel(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	util.WriteToExcel(user)

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processed task")
	return nil
}
