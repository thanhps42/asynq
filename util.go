package asynq

import "github.com/thanhps42/asynq/internal/base"

func DecodeMessage(data []byte) (*base.TaskMessage, error) {
	return base.DecodeMessage(data)
}
