package api

import (
	"context"
	"fmt"

	"quizzer/api/restapi"
	"quizzer/config"
)

type Handler struct {
	Tasks []config.Task
}

func (h Handler) APITaskIDGet(_ /* ctx */ context.Context, params restapi.APITaskIDGetParams) (*restapi.Task, error) {
	if params.ID >= uint64(len(h.Tasks)) {
		return nil, fmt.Errorf("invalid task id: %d", params.ID)
	}

	task := &restapi.Task{}
	task.Question = restapi.NewOptString(h.Tasks[params.ID].Question)

	for k, v := range h.Tasks[params.ID].Answers {
		task.Answers = append(task.Answers, restapi.Answer{
			ID:   restapi.NewOptUint64(uint64(k)),
			Text: restapi.NewOptString(v.Text),
		})
	}

	return task, nil
}

func (h Handler) APITaskIDResolveGet(_ /* ctx */ context.Context, params restapi.APITaskIDResolveGetParams) ([]restapi.Answer, error) {
	if params.ID >= uint64(len(h.Tasks)) {
		return nil, fmt.Errorf("invalid task id: %d", params.ID)
	}

	var answers []restapi.Answer

	for k, v := range h.Tasks[params.ID].Answers {
		if v.IsCorrect {
			answers = append(answers, restapi.Answer{
				ID:   restapi.NewOptUint64(uint64(k)),
				Text: restapi.NewOptString(v.Text),
			})
		}
	}

	return answers, nil
}

func (h Handler) APITaskIDSolveGet(_ /* ctx */ context.Context, params restapi.APITaskIDSolveGetParams) (bool, error) {
	if params.ID >= uint64(len(h.Tasks)) {
		return false, fmt.Errorf("invalid task id: %d", params.ID)
	}

	t := h.Tasks[params.ID]

	if params.SolutionID >= uint64(len(t.Answers)) {
		return false, fmt.Errorf("invalid answer id: %d", params.SolutionID)
	}

	return t.Answers[params.SolutionID].IsCorrect, nil
}
