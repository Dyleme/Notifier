package handler

import (
	"net/http"

	"github.com/Dyleme/Notifier/internal/authorization/authmiddleware"
	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/internal/service/handler/api"
	"github.com/Dyleme/Notifier/pkg/http/requests"
	"github.com/Dyleme/Notifier/pkg/http/responses"
	"github.com/Dyleme/Notifier/pkg/utils"
)

func (t EventHandler) ListTasks(w http.ResponseWriter, r *http.Request, params api.ListTasksParams) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	listParams := parseListParams(params.Offset, params.Limit)

	tasks, err := t.serv.ListUserTasks(r.Context(), userID, listParams)
	if err != nil {
		responses.KnownError(w, err)

		return
	}
	apiTasks := utils.DtoSlice(tasks, mapAPITask)

	responses.JSON(w, http.StatusOK, apiTasks)
}

func (t EventHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	var addTaskBody api.AddTaskJSONRequestBody
	err = requests.Bind(r, &addTaskBody)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)

		return
	}

	task := mapAddTaskReq(addTaskBody, userID)
	createdTask, err := t.serv.AddTask(r.Context(), task)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	responses.JSON(w, http.StatusCreated, mapAPITask(createdTask))
}

func mapAddTaskReq(body api.AddTaskJSONRequestBody, userID int) domains.Task {
	return domains.Task{ //nolint:exhaustruct // TODO: use separate struct for creation
		UserID: userID,
		Text:   body.Message,
	}
}

func mapUpdateTaskReq(body api.UpdateTaskReqBody, taskID, userID int) domains.Task {
	return domains.Task{
		ID:       taskID,
		UserID:   userID,
		Text:     body.Message,
		Periodic: body.Periodic,
		Archived: body.Archived,
	}
}

func mapAPITask(task domains.Task) api.Task {
	return api.Task{
		Id:       task.ID,
		Message:  task.Text,
		Archived: task.Archived,
		Periodic: task.Periodic,
	}
}

func (t EventHandler) GetTask(w http.ResponseWriter, r *http.Request, taskID int) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	task, err := t.serv.GetTask(r.Context(), taskID, userID)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	responses.JSON(w, http.StatusOK, mapAPITask(task))
}

func (t EventHandler) UpdateTask(w http.ResponseWriter, r *http.Request, taskID int) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	var reqBody api.UpdateTaskJSONRequestBody
	err = requests.Bind(r, &reqBody)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)

		return
	}

	task := mapUpdateTaskReq(reqBody, userID, taskID)

	err = t.serv.UpdateTask(r.Context(), task)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	responses.Status(w, http.StatusOK)
}
