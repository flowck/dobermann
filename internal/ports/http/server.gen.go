// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package http

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// CreateAccountRequest defines model for CreateAccountRequest.
type CreateAccountRequest struct {
	AccountName string `json:"account_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

// CreateMonitorRequest defines model for CreateMonitorRequest.
type CreateMonitorRequest struct {
	CheckIntervalInSeconds int    `json:"check_interval_in_seconds"`
	EndpointUrl            string `json:"endpoint_url"`
}

// EditMonitorRequest defines model for EditMonitorRequest.
type EditMonitorRequest struct {
	CheckIntervalInSeconds int    `json:"check_interval_in_seconds"`
	EndpointUrl            string `json:"endpoint_url"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error custom error code such as 'email_in_use'
	Error string `json:"error"`

	// Message A description about the error
	Message string `json:"message"`
}

// FullIncident defines model for FullIncident.
type FullIncident struct {
	Cause           string     `json:"cause"`
	CheckedUrl      string     `json:"checked_url"`
	CreatedAt       time.Time  `json:"created_at"`
	Id              string     `json:"id"`
	MonitorId       string     `json:"monitor_id"`
	RequestHeaders  string     `json:"request_headers"`
	ResolvedAt      *time.Time `json:"resolved_at,omitempty"`
	ResponseBody    string     `json:"response_body"`
	ResponseHeaders string     `json:"response_headers"`
	ResponseStatus  int        `json:"response_status"`
}

// GetAllIncidentsPayload defines model for GetAllIncidentsPayload.
type GetAllIncidentsPayload struct {
	Data       []Incident `json:"data"`
	Page       int        `json:"page"`
	PageCount  int        `json:"page_count"`
	PerPage    int        `json:"per_page"`
	TotalCount int64      `json:"total_count"`
}

// GetAllMonitorByIdPayload defines model for GetAllMonitorByIdPayload.
type GetAllMonitorByIdPayload struct {
	Data Monitor `json:"data"`
}

// GetAllMonitorsPayload defines model for GetAllMonitorsPayload.
type GetAllMonitorsPayload struct {
	Data       []Monitor `json:"data"`
	Page       int       `json:"page"`
	PageCount  int       `json:"page_count"`
	PerPage    int       `json:"per_page"`
	TotalCount int64     `json:"total_count"`
}

// GetIncidentByByIdPayload defines model for GetIncidentByByIdPayload.
type GetIncidentByByIdPayload struct {
	Data FullIncident `json:"data"`
}

// GetMonitorResponseTimeStatsPayload defines model for GetMonitorResponseTimeStatsPayload.
type GetMonitorResponseTimeStatsPayload struct {
	Data []ResponseTimeStat `json:"data"`
}

// GetProfileDetailsPayload defines model for GetProfileDetailsPayload.
type GetProfileDetailsPayload struct {
	Data User `json:"data"`
}

// Incident defines model for Incident.
type Incident struct {
	Cause      string     `json:"cause"`
	CheckedUrl string     `json:"checked_url"`
	CreatedAt  time.Time  `json:"created_at"`
	Id         string     `json:"id"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
}

// LogInPayload defines model for LogInPayload.
type LogInPayload struct {
	Token string `json:"token"`
}

// LogInRequest defines model for LogInRequest.
type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Monitor defines model for Monitor.
type Monitor struct {
	CheckIntervalInSeconds int        `json:"check_interval_in_seconds"`
	CreatedAt              time.Time  `json:"created_at"`
	DownSince              *time.Time `json:"down_since,omitempty"`
	EndpointUrl            string     `json:"endpoint_url"`
	Id                     string     `json:"id"`
	Incidents              []Incident `json:"incidents"`
	IsEndpointUp           bool       `json:"is_endpoint_up"`
	IsPaused               bool       `json:"is_paused"`
	LastCheckedAt          *time.Time `json:"last_checked_at,omitempty"`
	UpSince                *time.Time `json:"up_since,omitempty"`
}

// ResponseTimeStat defines model for ResponseTimeStat.
type ResponseTimeStat struct {
	Date   time.Time `json:"date"`
	Region string    `json:"region"`
	Value  int       `json:"value"`
}

// ToggleMonitorPauseRequest defines model for ToggleMonitorPauseRequest.
type ToggleMonitorPauseRequest struct {
	Pause bool `json:"pause"`
}

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	Id        string    `json:"id"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
}

// DefaultError defines model for DefaultError.
type DefaultError = ErrorResponse

// GetAllIncidentsParams defines parameters for GetAllIncidents.
type GetAllIncidentsParams struct {
	Page  *int `form:"page,omitempty" json:"page,omitempty"`
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetAllMonitorsParams defines parameters for GetAllMonitors.
type GetAllMonitorsParams struct {
	Page  *int `form:"page,omitempty" json:"page,omitempty"`
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetMonitorResponseTimeStatsParams defines parameters for GetMonitorResponseTimeStats.
type GetMonitorResponseTimeStatsParams struct {
	RangeInDays *int `form:"range_in_days,omitempty" json:"range_in_days,omitempty"`
}

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody = CreateAccountRequest

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LogInRequest

// CreateMonitorJSONRequestBody defines body for CreateMonitor for application/json ContentType.
type CreateMonitorJSONRequestBody = CreateMonitorRequest

// ToggleMonitorPauseJSONRequestBody defines body for ToggleMonitorPause for application/json ContentType.
type ToggleMonitorPauseJSONRequestBody = ToggleMonitorPauseRequest

// EditMonitorJSONRequestBody defines body for EditMonitor for application/json ContentType.
type EditMonitorJSONRequestBody = EditMonitorRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get details about the user currently logged in
	// (GET /accounts/profile)
	GetProfileDetails(ctx echo.Context) error
	// Creates a new account
	// (POST /auth/accounts)
	CreateAccount(ctx echo.Context) error
	// Log in
	// (POST /auth/login)
	Login(ctx echo.Context) error
	// Get all incidents
	// (GET /incidents)
	GetAllIncidents(ctx echo.Context, params GetAllIncidentsParams) error
	// Get an incident by id
	// (GET /incidents/{incidentID})
	GetIncidentByID(ctx echo.Context, incidentID string) error
	// Get all monitors in a with pagination
	// (GET /monitors)
	GetAllMonitors(ctx echo.Context, params GetAllMonitorsParams) error
	// Create a new monitor
	// (POST /monitors)
	CreateMonitor(ctx echo.Context) error
	// Delete monitor
	// (DELETE /monitors/{monitorID})
	DeleteMonitor(ctx echo.Context, monitorID string) error
	// Get all monitors in a with pagination
	// (GET /monitors/{monitorID})
	GetMonitorByID(ctx echo.Context, monitorID string) error
	// Pause or unpause the monitor
	// (POST /monitors/{monitorID})
	ToggleMonitorPause(ctx echo.Context, monitorID string) error
	// Edit a monitor by id
	// (PUT /monitors/{monitorID})
	EditMonitor(ctx echo.Context, monitorID string) error
	// Get the stats about the response time
	// (GET /monitors/{monitorID}/stats/response-times)
	GetMonitorResponseTimeStats(ctx echo.Context, monitorID string, params GetMonitorResponseTimeStatsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetProfileDetails converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfileDetails(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfileDetails(ctx)
	return err
}

// CreateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAccount(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateAccount(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// GetAllIncidents converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllIncidents(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllIncidentsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllIncidents(ctx, params)
	return err
}

// GetIncidentByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetIncidentByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentID" -------------
	var incidentID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentID", runtime.ParamLocationPath, ctx.Param("incidentID"), &incidentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetIncidentByID(ctx, incidentID)
	return err
}

// GetAllMonitors converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllMonitors(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllMonitorsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllMonitors(ctx, params)
	return err
}

// CreateMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) CreateMonitor(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateMonitor(ctx)
	return err
}

// DeleteMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMonitor(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "monitorID" -------------
	var monitorID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, ctx.Param("monitorID"), &monitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter monitorID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteMonitor(ctx, monitorID)
	return err
}

// GetMonitorByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetMonitorByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "monitorID" -------------
	var monitorID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, ctx.Param("monitorID"), &monitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter monitorID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMonitorByID(ctx, monitorID)
	return err
}

// ToggleMonitorPause converts echo context to params.
func (w *ServerInterfaceWrapper) ToggleMonitorPause(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "monitorID" -------------
	var monitorID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, ctx.Param("monitorID"), &monitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter monitorID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ToggleMonitorPause(ctx, monitorID)
	return err
}

// EditMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) EditMonitor(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "monitorID" -------------
	var monitorID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, ctx.Param("monitorID"), &monitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter monitorID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.EditMonitor(ctx, monitorID)
	return err
}

// GetMonitorResponseTimeStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetMonitorResponseTimeStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "monitorID" -------------
	var monitorID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, ctx.Param("monitorID"), &monitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter monitorID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMonitorResponseTimeStatsParams
	// ------------- Optional query parameter "range_in_days" -------------

	err = runtime.BindQueryParameter("form", true, false, "range_in_days", ctx.QueryParams(), &params.RangeInDays)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter range_in_days: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMonitorResponseTimeStats(ctx, monitorID, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/accounts/profile", wrapper.GetProfileDetails)
	router.POST(baseURL+"/auth/accounts", wrapper.CreateAccount)
	router.POST(baseURL+"/auth/login", wrapper.Login)
	router.GET(baseURL+"/incidents", wrapper.GetAllIncidents)
	router.GET(baseURL+"/incidents/:incidentID", wrapper.GetIncidentByID)
	router.GET(baseURL+"/monitors", wrapper.GetAllMonitors)
	router.POST(baseURL+"/monitors", wrapper.CreateMonitor)
	router.DELETE(baseURL+"/monitors/:monitorID", wrapper.DeleteMonitor)
	router.GET(baseURL+"/monitors/:monitorID", wrapper.GetMonitorByID)
	router.POST(baseURL+"/monitors/:monitorID", wrapper.ToggleMonitorPause)
	router.PUT(baseURL+"/monitors/:monitorID", wrapper.EditMonitor)
	router.GET(baseURL+"/monitors/:monitorID/stats/response-times", wrapper.GetMonitorResponseTimeStats)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"Xm//DAAA//+E9ncKLiYAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
