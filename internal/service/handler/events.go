package handler

import (
	"net/http"
	"time"

	"github.com/Dyleme/Notifier/internal/authorization/authmiddleware"
	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/internal/service/handler/api"
	"github.com/Dyleme/Notifier/internal/service/service"
	"github.com/Dyleme/Notifier/pkg/http/requests"
	"github.com/Dyleme/Notifier/pkg/http/responses"
)

func (t EventHandler) ListEvents(w http.ResponseWriter, r *http.Request, params api.ListEventsParams) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}
	var from, to time.Time

	if params.From == nil {
		from = time.Time{}
	}
	if params.To == nil {
		to = time.Now().Add(30 * 24 * time.Hour)
	}

	listParams := parseListParams(params.Offset, params.Limit)

	tasks, err := t.serv.ListEventsInPeriod(r.Context(), userID, from, to, listParams)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	responses.JSON(w, http.StatusOK, tasks)
}

func (t EventHandler) PostEventSetTaskID(w http.ResponseWriter, r *http.Request, taskID int) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	var setEvent api.SetEventReqBody
	err = requests.Bind(r, &setEvent)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)

		return
	}
	description := ""
	if setEvent.Description != nil {
		description = *setEvent.Description
	}

	event, err := t.serv.AddTaskToEvent(r.Context(), userID, taskID, setEvent.Start, description)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	apiEvent := mapAPIEvent(event)

	responses.JSON(w, http.StatusOK, apiEvent)
}

func mapAPIEvent(event domains.Event) api.Event {
	return api.Event{
		Description: &event.Description,
		Done:        event.Done,
		Id:          event.ID,
		Start:       event.Start,
		Text:        event.Text,
	}
}

func (t EventHandler) GetEvent(w http.ResponseWriter, r *http.Request, eventID int) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	event, err := t.serv.GetEvent(r.Context(), userID, eventID)
	if err != nil {
		responses.KnownError(w, err)

		return
	}

	apiEvent := mapAPIEvent(event)

	responses.JSON(w, http.StatusOK, apiEvent)
}

func (t EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request, eventID int) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	var updateBody api.UpdateEventJSONRequestBody
	err = requests.Bind(r, &updateBody)
	if err != nil {
		responses.KnownError(w, err)

		return
	}
	description := ""
	if updateBody.Description != nil {
		description = *updateBody.Description
	}

	event, err := t.serv.UpdateEvent(r.Context(), service.EventUpdateParams{
		ID:          eventID,
		UserID:      userID,
		Text:        description,
		Description: description,
		Start:       updateBody.Start,
	})
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	apiEvent := mapAPIEvent(event)

	responses.JSON(w, http.StatusOK, apiEvent)
}

func mapCreateEvent(body api.CreateEventReqBody, userID int) domains.Event {
	description := ""
	if body.Description != nil {
		description = *body.Description
	}

	event := domains.Event{ //nolint:exhaustruct //creation object we don't know ids
		UserID:             userID,
		Text:               body.Message,
		Description:        description,
		Start:              body.Start,
		Done:               false,
		Sended:             false,
		NotificationParams: nil,
		SendTime:           body.Start,
	}

	return event
}

func (t EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := authmiddleware.UserIDFromCtx(r.Context())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	var createBody api.CreateEventReqBody
	err = requests.Bind(r, &createBody)
	if err != nil {
		responses.KnownError(w, err)

		return
	}
	event := mapCreateEvent(createBody, userID)
	createdEvent, err := t.serv.CreateEvent(r.Context(), event)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)

		return
	}

	apiEvent := mapAPIEvent(createdEvent)

	responses.JSON(w, http.StatusOK, apiEvent)
}
