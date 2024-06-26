// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreateTaskReqBody defines model for CreateTaskReqBody.
type CreateTaskReqBody struct {
	Description *string   `json:"description,omitempty"`
	Message     string    `json:"message"`
	Start       time.Time `json:"start"`
}

// EventInfo defines model for EventInfo.
type EventInfo struct {
	Cmd      *bool   `json:"cmd,omitempty"`
	Telegram *int    `json:"telegram,omitempty"`
	Webhook  *string `json:"webhook,omitempty"`
}

// EventParams defines model for EventParams.
type EventParams struct {
	DelayedTill *time.Time `json:"delayedTill,omitempty"`
	Info        EventInfo  `json:"info"`

	// Period Required time for task in minutes
	Period int `json:"period"`
}

// Task defines model for Task.
type Task struct {
	Description *string   `json:"description,omitempty"`
	Done        bool      `json:"done"`
	Id          int       `json:"id"`
	Start       time.Time `json:"start"`
	Text        string    `json:"text"`
}

// UpdateTaskReqBody defines model for UpdateTaskReqBody.
type UpdateTaskReqBody struct {
	Description *string   `json:"description,omitempty"`
	Done        bool      `json:"done"`
	Start       time.Time `json:"start"`
}

// LimitParam defines model for limitParam.
type LimitParam = int32

// OffsetParam defines model for offsetParam.
type OffsetParam = int32

// ListTasksParams defines parameters for ListTasks.
type ListTasksParams struct {
	From *time.Time `form:"from,omitempty" json:"from,omitempty"`
	To   *time.Time `form:"to,omitempty" json:"to,omitempty"`

	// Limit Limits the number of returned results
	Limit *LimitParam `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset from which start returned results
	Offset *OffsetParam `form:"offset,omitempty" json:"offset,omitempty"`
}

// SetDefaultEventParamsJSONRequestBody defines body for SetDefaultEventParams for application/json ContentType.
type SetDefaultEventParamsJSONRequestBody = EventParams

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = CreateTaskReqBody

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody = UpdateTaskReqBody

// SetTaskEventParamsJSONRequestBody defines body for SetTaskEventParams for application/json ContentType.
type SetTaskEventParamsJSONRequestBody = EventParams

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Set default event params
	// (GET /event-params)
	GetDefaultEventParams(w http.ResponseWriter, r *http.Request)
	// Set default event params
	// (PUT /event-params)
	SetDefaultEventParams(w http.ResponseWriter, r *http.Request)
	// List tasks
	// (GET /task)
	ListTasks(w http.ResponseWriter, r *http.Request, params ListTasksParams)
	// Create new task
	// (POST /task)
	CreateTask(w http.ResponseWriter, r *http.Request)
	// Get task
	// (GET /task/{taskID})
	GetTask(w http.ResponseWriter, r *http.Request, taskID int)
	// Update task
	// (PUT /task/{taskID})
	UpdateTask(w http.ResponseWriter, r *http.Request, taskID int)
	// Get task event params
	// (GET /task/{taskID}/event-params)
	GetTaskEventParams(w http.ResponseWriter, r *http.Request, taskID int)
	// Set Task event params
	// (PUT /task/{taskID}/event-params)
	SetTaskEventParams(w http.ResponseWriter, r *http.Request, taskID int)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Set default event params
// (GET /event-params)
func (_ Unimplemented) GetDefaultEventParams(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Set default event params
// (PUT /event-params)
func (_ Unimplemented) SetDefaultEventParams(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List tasks
// (GET /task)
func (_ Unimplemented) ListTasks(w http.ResponseWriter, r *http.Request, params ListTasksParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create new task
// (POST /task)
func (_ Unimplemented) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get task
// (GET /task/{taskID})
func (_ Unimplemented) GetTask(w http.ResponseWriter, r *http.Request, taskID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update task
// (PUT /task/{taskID})
func (_ Unimplemented) UpdateTask(w http.ResponseWriter, r *http.Request, taskID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get task event params
// (GET /task/{taskID}/event-params)
func (_ Unimplemented) GetTaskEventParams(w http.ResponseWriter, r *http.Request, taskID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Set Task event params
// (PUT /task/{taskID}/event-params)
func (_ Unimplemented) SetTaskEventParams(w http.ResponseWriter, r *http.Request, taskID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetDefaultEventParams operation middleware
func (siw *ServerInterfaceWrapper) GetDefaultEventParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDefaultEventParams(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SetDefaultEventParams operation middleware
func (siw *ServerInterfaceWrapper) SetDefaultEventParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetDefaultEventParams(w, r)
	}))

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

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListTasks(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r)
	}))

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

	err = runtime.BindStyledParameterWithOptions("simple", "taskID", chi.URLParam(r, "taskID"), &taskID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTask(w, r, taskID)
	}))

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

	err = runtime.BindStyledParameterWithOptions("simple", "taskID", chi.URLParam(r, "taskID"), &taskID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateTask(w, r, taskID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTaskEventParams operation middleware
func (siw *ServerInterfaceWrapper) GetTaskEventParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID int

	err = runtime.BindStyledParameterWithOptions("simple", "taskID", chi.URLParam(r, "taskID"), &taskID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTaskEventParams(w, r, taskID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SetTaskEventParams operation middleware
func (siw *ServerInterfaceWrapper) SetTaskEventParams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "taskID" -------------
	var taskID int

	err = runtime.BindStyledParameterWithOptions("simple", "taskID", chi.URLParam(r, "taskID"), &taskID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "taskID", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetTaskEventParams(w, r, taskID)
	}))

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

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
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
		r.Get(options.BaseURL+"/event-params", wrapper.GetDefaultEventParams)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/event-params", wrapper.SetDefaultEventParams)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/task", wrapper.ListTasks)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/task", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/task/{taskID}", wrapper.GetTask)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/task/{taskID}", wrapper.UpdateTask)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/task/{taskID}/event-params", wrapper.GetTaskEventParams)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/task/{taskID}/event-params", wrapper.SetTaskEventParams)
	})

	return r
}
