package api

import (
	"quizzer/api/models"
	"quizzer/api/restapi/operations"
	"quizzer/config"

	"github.com/go-openapi/runtime/middleware"
)

func GetTask(tasks []config.Task) func(operations.GetAPITaskIDParams) middleware.Responder {
	return func(params operations.GetAPITaskIDParams) middleware.Responder {
		if params.ID >= uint64(len(tasks)) {
			return operations.NewGetAPITaskIDOK().WithPayload(nil)
		}

		task := models.Task{}
		task.Question = tasks[params.ID].Question
		task.Answers = []*models.Answer{}

		for k, v := range tasks[params.ID].Answers {
			task.Answers = append(task.Answers, &models.Answer{ID: uint64(k), Text: v.Text})
		}

		return operations.NewGetAPITaskIDOK().WithPayload(&task)
	}
}
