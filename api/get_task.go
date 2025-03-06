package api

import (
	"context"
	"fmt"

	"quizzer/api/restapi"
)

func (h Handler) APITaskIDGet(ctx context.Context, params restapi.APITaskIDGetParams) (*restapi.Task, error) {
	if params.ID >= uint64(len(h.Tasks)) {
		return nil, fmt.Errorf("invalid task id: %d", params.ID)
	}

	task := &restapi.Task{}
	task.Question = restapi.NewOptString(h.Tasks[params.ID].Question)
	task.Answers = []restapi.Answer{}

	for k, v := range h.Tasks[params.ID].Answers {
		task.Answers = append(task.Answers, restapi.Answer{
			ID:   restapi.NewOptUint64(uint64(k)),
			Text: restapi.NewOptString(v.Text),
		})
	}

	return task, nil
}
