package api

import (
	"context"

	"quizzer/api/restapi"
	"quizzer/config"
)

type Handler struct {
	Tasks []config.Task
}

func (h Handler) APITaskIDResolveGet(ctx context.Context, params restapi.APITaskIDResolveGetParams) ([]restapi.Answer, error) {
	// TODO implement me
	panic("implement me")
}

func (h Handler) APITaskIDSolveGet(ctx context.Context, params restapi.APITaskIDSolveGetParams) (bool, error) {
	// TODO implement me
	panic("implement me")
}
