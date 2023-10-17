// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreateEventReqBody defines model for CreateEventReqBody.
type CreateEventReqBody struct {
	Description *string   `json:"description,omitempty"`
	Message     string    `json:"message"`
	Start       time.Time `json:"start"`
}

// CreateTaskReqBody defines model for CreateTaskReqBody.
type CreateTaskReqBody struct {
	Message  string `json:"message"`
	Periodic bool   `json:"periodic"`
}

// Event defines model for Event.
type Event struct {
	Description *string   `json:"description,omitempty"`
	Done        bool      `json:"done"`
	Id          int       `json:"id"`
	Start       time.Time `json:"start"`
	Text        string    `json:"text"`
}

// NotificationInfo defines model for NotificationInfo.
type NotificationInfo struct {
	Cmd      *bool   `json:"cmd,omitempty"`
	Telegram *int    `json:"telegram,omitempty"`
	Webhook  *string `json:"webhook,omitempty"`
}

// NotificationParams defines model for NotificationParams.
type NotificationParams struct {
	DelayedTill *time.Time       `json:"delayedTill,omitempty"`
	Info        NotificationInfo `json:"info"`

	// Period Required time for task in minutes
	Period int `json:"period"`
}

// SetEventReqBody defines model for SetEventReqBody.
type SetEventReqBody struct {
	Description *string   `json:"description,omitempty"`
	Start       time.Time `json:"start"`
}

// Task defines model for Task.
type Task struct {
	Archived bool   `json:"archived"`
	Id       int    `json:"id"`
	Message  string `json:"message"`
	Periodic bool   `json:"periodic"`
}

// UpdateEventReqBody defines model for UpdateEventReqBody.
type UpdateEventReqBody struct {
	Description *string   `json:"description,omitempty"`
	Done        bool      `json:"done"`
	Start       time.Time `json:"start"`
}

// UpdateTaskReqBody defines model for UpdateTaskReqBody.
type UpdateTaskReqBody struct {
	Archived bool   `json:"archived"`
	Done     bool   `json:"done"`
	Message  string `json:"message"`
	Periodic bool   `json:"periodic"`

	// RequiredTime Required time for task in minutes
	RequiredTime int `json:"requiredTime"`
}

// LimitParam defines model for limitParam.
type LimitParam = int32

// OffsetParam defines model for offsetParam.
type OffsetParam = int32

// ListEventsParams defines parameters for ListEvents.
type ListEventsParams struct {
	From *time.Time `form:"from,omitempty" json:"from,omitempty"`
	To   *time.Time `form:"to,omitempty" json:"to,omitempty"`

	// Limit Limits the number of returned results
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset from which start returned results
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`
}

// ListTasksParams defines parameters for ListTasks.
type ListTasksParams struct {
	// Limit Limits the number of returned results
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset from which start returned results
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`
}

// CreateEventJSONRequestBody defines body for CreateEvent for application/json ContentType.
type CreateEventJSONRequestBody = CreateEventReqBody

// PostEventSetTaskIDJSONRequestBody defines body for PostEventSetTaskID for application/json ContentType.
type PostEventSetTaskIDJSONRequestBody = SetEventReqBody

// UpdateEventJSONRequestBody defines body for UpdateEvent for application/json ContentType.
type UpdateEventJSONRequestBody = UpdateEventReqBody

// SetEventNotificationParamsJSONRequestBody defines body for SetEventNotificationParams for application/json ContentType.
type SetEventNotificationParamsJSONRequestBody = NotificationParams

// SetDefaultNotificationParamsJSONRequestBody defines body for SetDefaultNotificationParams for application/json ContentType.
type SetDefaultNotificationParamsJSONRequestBody = NotificationParams

// AddTaskJSONRequestBody defines body for AddTask for application/json ContentType.
type AddTaskJSONRequestBody = CreateTaskReqBody

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody = UpdateTaskReqBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List events
	// (GET /event)
	ListEvents(w http.ResponseWriter, r *http.Request, params ListEventsParams)
	// Create new event
	// (POST /event)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	// Add a new event
	// (POST /event/set/{taskID})
	PostEventSetTaskID(w http.ResponseWriter, r *http.Request, taskID int)
	// Get event
	// (GET /event/{eventID})
	GetEvent(w http.ResponseWriter, r *http.Request, eventID int)
	// Update event
	// (PUT /event/{eventID})
	UpdateEvent(w http.ResponseWriter, r *http.Request, eventID int)
	// Get event notification params
	// (GET /event/{eventID}/notification-params)
	GetEventNotificationParams(w http.ResponseWriter, r *http.Request, eventID int)
	// Set Event notification params
	// (PUT /event/{eventID}/notification-params)
	SetEventNotificationParams(w http.ResponseWriter, r *http.Request, eventID int)
	// Set default notification params
	// (GET /notification-params)
	GetDefaultNotificationParams(w http.ResponseWriter, r *http.Request)
	// Set default notification params
	// (PUT /notification-params)
	SetDefaultNotificationParams(w http.ResponseWriter, r *http.Request)
	// List user tasks
	// (GET /task)
	ListTasks(w http.ResponseWriter, r *http.Request, params ListTasksParams)
	// Add a new task
	// (POST /task)
	AddTask(w http.ResponseWriter, r *http.Request)
	// Get task
	// (GET /task/{taskID})
	GetTask(w http.ResponseWriter, r *http.Request, taskID int)
	// Add a new task
	// (PUT /task/{taskID})
	UpdateTask(w http.ResponseWriter, r *http.Request, taskID int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListEvents operation middleware
func (siw *ServerInterfaceWrapper) ListEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListEventsParams

	// ------------- Optional query parameter "from" -------------

	err = runtime.BindQueryParameter("form", true, false, "from", r.URL.Query(), &params.From)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "from", Err: err})
		return
	}

	// ------------- Optional query parameter "to" -------------

	err = runtime.BindQueryParameter("form", true, false, "to", r.URL.Query(), &params.To)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "to", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListEvents(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateEvent operation middleware
func (siw *ServerInterfaceWrapper) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateEvent(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostEventSetTaskID operation middleware
func (siw *ServerInterfaceWrapper) PostEventSetTaskID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostEventSetTaskID(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetEvent operation middleware
func (siw *ServerInterfaceWrapper) GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "eventID" -------------
	var eventID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventID", runtime.ParamLocationPath, chi.URLParam(r, "eventID"), &eventID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "eventID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetEvent(w, r, eventID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateEvent operation middleware
func (siw *ServerInterfaceWrapper) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "eventID" -------------
	var eventID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventID", runtime.ParamLocationPath, chi.URLParam(r, "eventID"), &eventID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "eventID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateEvent(w, r, eventID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetEventNotificationParams operation middleware
func (siw *ServerInterfaceWrapper) GetEventNotificationParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "eventID" -------------
	var eventID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventID", runtime.ParamLocationPath, chi.URLParam(r, "eventID"), &eventID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "eventID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetEventNotificationParams(w, r, eventID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SetEventNotificationParams operation middleware
func (siw *ServerInterfaceWrapper) SetEventNotificationParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "eventID" -------------
	var eventID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "eventID", runtime.ParamLocationPath, chi.URLParam(r, "eventID"), &eventID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "eventID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetEventNotificationParams(w, r, eventID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetDefaultNotificationParams operation middleware
func (siw *ServerInterfaceWrapper) GetDefaultNotificationParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDefaultNotificationParams(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SetDefaultNotificationParams operation middleware
func (siw *ServerInterfaceWrapper) SetDefaultNotificationParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetDefaultNotificationParams(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListTasks operation middleware
func (siw *ServerInterfaceWrapper) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListTasksParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListTasks(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddTask operation middleware
func (siw *ServerInterfaceWrapper) AddTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddTask(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTask operation middleware
func (siw *ServerInterfaceWrapper) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateTask operation middleware
func (siw *ServerInterfaceWrapper) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID int

	err = runtime.BindStyledParameterWithLocation("simple", false, "taskID", runtime.ParamLocationPath, chi.URLParam(r, "taskID"), &taskID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateTask(w, r, taskID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/event", wrapper.ListEvents)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/event", wrapper.CreateEvent)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/event/set/{taskID}", wrapper.PostEventSetTaskID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/event/{eventID}", wrapper.GetEvent)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/event/{eventID}", wrapper.UpdateEvent)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/event/{eventID}/notification-params", wrapper.GetEventNotificationParams)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/event/{eventID}/notification-params", wrapper.SetEventNotificationParams)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/notification-params", wrapper.GetDefaultNotificationParams)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/notification-params", wrapper.SetDefaultNotificationParams)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/task", wrapper.ListTasks)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/task", wrapper.AddTask)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/task/{taskID}", wrapper.GetTask)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/task/{taskID}", wrapper.UpdateTask)
	})

	return r
}