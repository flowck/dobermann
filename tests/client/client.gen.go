// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
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

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error custom error code such as 'email_in_use'
	Error string `json:"error"`

	// Message A description about the error
	Message string `json:"message"`
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

// Incident defines model for Incident.
type Incident struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
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
	EndpointUrl            string     `json:"endpoint_url"`
	Id                     string     `json:"id"`
	Incidents              []Incident `json:"incidents"`
	IsEndpointUp           bool       `json:"is_endpoint_up"`
	IsPaused               bool       `json:"is_paused"`
	LastCheckedAt          *time.Time `json:"last_checked_at,omitempty"`
}

// ToggleMonitorPauseRequest defines model for ToggleMonitorPauseRequest.
type ToggleMonitorPauseRequest struct {
	Pause bool `json:"pause"`
}

// DefaultError defines model for DefaultError.
type DefaultError = ErrorResponse

// GetAllMonitorsParams defines parameters for GetAllMonitors.
type GetAllMonitorsParams struct {
	Page  *int `form:"page,omitempty" json:"page,omitempty"`
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody = CreateAccountRequest

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LogInRequest

// CreateMonitorJSONRequestBody defines body for CreateMonitor for application/json ContentType.
type CreateMonitorJSONRequestBody = CreateMonitorRequest

// ToggleMonitorPauseJSONRequestBody defines body for ToggleMonitorPause for application/json ContentType.
type ToggleMonitorPauseJSONRequestBody = ToggleMonitorPauseRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateAccount request with any body
	CreateAccountWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAccount(ctx context.Context, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Login request with any body
	LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Login(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAllMonitors request
	GetAllMonitors(ctx context.Context, params *GetAllMonitorsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateMonitor request with any body
	CreateMonitorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateMonitor(ctx context.Context, body CreateMonitorJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteMonitor request
	DeleteMonitor(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetMonitorByID request
	GetMonitorByID(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ToggleMonitorPause request with any body
	ToggleMonitorPauseWithBody(ctx context.Context, monitorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ToggleMonitorPause(ctx context.Context, monitorID string, body ToggleMonitorPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateAccountWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAccountRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAccount(ctx context.Context, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAccountRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Login(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAllMonitors(ctx context.Context, params *GetAllMonitorsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllMonitorsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateMonitorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateMonitorRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateMonitor(ctx context.Context, body CreateMonitorJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateMonitorRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteMonitor(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteMonitorRequest(c.Server, monitorID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetMonitorByID(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMonitorByIDRequest(c.Server, monitorID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ToggleMonitorPauseWithBody(ctx context.Context, monitorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewToggleMonitorPauseRequestWithBody(c.Server, monitorID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ToggleMonitorPause(ctx context.Context, monitorID string, body ToggleMonitorPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewToggleMonitorPauseRequest(c.Server, monitorID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateAccountRequest calls the generic CreateAccount builder with application/json body
func NewCreateAccountRequest(server string, body CreateAccountJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAccountRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateAccountRequestWithBody generates requests for CreateAccount with any type of body
func NewCreateAccountRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/accounts")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewLoginRequest calls the generic Login builder with application/json body
func NewLoginRequest(server string, body LoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewLoginRequestWithBody(server, "application/json", bodyReader)
}

// NewLoginRequestWithBody generates requests for Login with any type of body
func NewLoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/auth/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetAllMonitorsRequest generates requests for GetAllMonitors
func NewGetAllMonitorsRequest(server string, params *GetAllMonitorsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/monitors")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateMonitorRequest calls the generic CreateMonitor builder with application/json body
func NewCreateMonitorRequest(server string, body CreateMonitorJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateMonitorRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateMonitorRequestWithBody generates requests for CreateMonitor with any type of body
func NewCreateMonitorRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/monitors")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteMonitorRequest generates requests for DeleteMonitor
func NewDeleteMonitorRequest(server string, monitorID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, monitorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/monitors/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetMonitorByIDRequest generates requests for GetMonitorByID
func NewGetMonitorByIDRequest(server string, monitorID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, monitorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/monitors/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewToggleMonitorPauseRequest calls the generic ToggleMonitorPause builder with application/json body
func NewToggleMonitorPauseRequest(server string, monitorID string, body ToggleMonitorPauseJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewToggleMonitorPauseRequestWithBody(server, monitorID, "application/json", bodyReader)
}

// NewToggleMonitorPauseRequestWithBody generates requests for ToggleMonitorPause with any type of body
func NewToggleMonitorPauseRequestWithBody(server string, monitorID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "monitorID", runtime.ParamLocationPath, monitorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/monitors/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateAccount request with any body
	CreateAccountWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error)

	CreateAccountWithResponse(ctx context.Context, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error)

	// Login request with any body
	LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginResponse, error)

	LoginWithResponse(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginResponse, error)

	// GetAllMonitors request
	GetAllMonitorsWithResponse(ctx context.Context, params *GetAllMonitorsParams, reqEditors ...RequestEditorFn) (*GetAllMonitorsResponse, error)

	// CreateMonitor request with any body
	CreateMonitorWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateMonitorResponse, error)

	CreateMonitorWithResponse(ctx context.Context, body CreateMonitorJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateMonitorResponse, error)

	// DeleteMonitor request
	DeleteMonitorWithResponse(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*DeleteMonitorResponse, error)

	// GetMonitorByID request
	GetMonitorByIDWithResponse(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*GetMonitorByIDResponse, error)

	// ToggleMonitorPause request with any body
	ToggleMonitorPauseWithBodyWithResponse(ctx context.Context, monitorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ToggleMonitorPauseResponse, error)

	ToggleMonitorPauseWithResponse(ctx context.Context, monitorID string, body ToggleMonitorPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*ToggleMonitorPauseResponse, error)
}

type CreateAccountResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateAccountResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAccountResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LogInPayload
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r LoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAllMonitorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAllMonitorsPayload
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetAllMonitorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllMonitorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateMonitorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateMonitorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateMonitorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteMonitorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeleteMonitorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteMonitorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMonitorByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetAllMonitorByIdPayload
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetMonitorByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetMonitorByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ToggleMonitorPauseResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r ToggleMonitorPauseResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ToggleMonitorPauseResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateAccountWithBodyWithResponse request with arbitrary body returning *CreateAccountResponse
func (c *ClientWithResponses) CreateAccountWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error) {
	rsp, err := c.CreateAccountWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAccountResponse(rsp)
}

func (c *ClientWithResponses) CreateAccountWithResponse(ctx context.Context, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error) {
	rsp, err := c.CreateAccount(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAccountResponse(rsp)
}

// LoginWithBodyWithResponse request with arbitrary body returning *LoginResponse
func (c *ClientWithResponses) LoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginResponse, error) {
	rsp, err := c.LoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginResponse(rsp)
}

func (c *ClientWithResponses) LoginWithResponse(ctx context.Context, body LoginJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginResponse, error) {
	rsp, err := c.Login(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginResponse(rsp)
}

// GetAllMonitorsWithResponse request returning *GetAllMonitorsResponse
func (c *ClientWithResponses) GetAllMonitorsWithResponse(ctx context.Context, params *GetAllMonitorsParams, reqEditors ...RequestEditorFn) (*GetAllMonitorsResponse, error) {
	rsp, err := c.GetAllMonitors(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllMonitorsResponse(rsp)
}

// CreateMonitorWithBodyWithResponse request with arbitrary body returning *CreateMonitorResponse
func (c *ClientWithResponses) CreateMonitorWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateMonitorResponse, error) {
	rsp, err := c.CreateMonitorWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateMonitorResponse(rsp)
}

func (c *ClientWithResponses) CreateMonitorWithResponse(ctx context.Context, body CreateMonitorJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateMonitorResponse, error) {
	rsp, err := c.CreateMonitor(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateMonitorResponse(rsp)
}

// DeleteMonitorWithResponse request returning *DeleteMonitorResponse
func (c *ClientWithResponses) DeleteMonitorWithResponse(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*DeleteMonitorResponse, error) {
	rsp, err := c.DeleteMonitor(ctx, monitorID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteMonitorResponse(rsp)
}

// GetMonitorByIDWithResponse request returning *GetMonitorByIDResponse
func (c *ClientWithResponses) GetMonitorByIDWithResponse(ctx context.Context, monitorID string, reqEditors ...RequestEditorFn) (*GetMonitorByIDResponse, error) {
	rsp, err := c.GetMonitorByID(ctx, monitorID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMonitorByIDResponse(rsp)
}

// ToggleMonitorPauseWithBodyWithResponse request with arbitrary body returning *ToggleMonitorPauseResponse
func (c *ClientWithResponses) ToggleMonitorPauseWithBodyWithResponse(ctx context.Context, monitorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ToggleMonitorPauseResponse, error) {
	rsp, err := c.ToggleMonitorPauseWithBody(ctx, monitorID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseToggleMonitorPauseResponse(rsp)
}

func (c *ClientWithResponses) ToggleMonitorPauseWithResponse(ctx context.Context, monitorID string, body ToggleMonitorPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*ToggleMonitorPauseResponse, error) {
	rsp, err := c.ToggleMonitorPause(ctx, monitorID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseToggleMonitorPauseResponse(rsp)
}

// ParseCreateAccountResponse parses an HTTP response from a CreateAccountWithResponse call
func ParseCreateAccountResponse(rsp *http.Response) (*CreateAccountResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAccountResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseLoginResponse parses an HTTP response from a LoginWithResponse call
func ParseLoginResponse(rsp *http.Response) (*LoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LogInPayload
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetAllMonitorsResponse parses an HTTP response from a GetAllMonitorsWithResponse call
func ParseGetAllMonitorsResponse(rsp *http.Response) (*GetAllMonitorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllMonitorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAllMonitorsPayload
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateMonitorResponse parses an HTTP response from a CreateMonitorWithResponse call
func ParseCreateMonitorResponse(rsp *http.Response) (*CreateMonitorResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateMonitorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeleteMonitorResponse parses an HTTP response from a DeleteMonitorWithResponse call
func ParseDeleteMonitorResponse(rsp *http.Response) (*DeleteMonitorResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteMonitorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetMonitorByIDResponse parses an HTTP response from a GetMonitorByIDWithResponse call
func ParseGetMonitorByIDResponse(rsp *http.Response) (*GetMonitorByIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetMonitorByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetAllMonitorByIdPayload
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseToggleMonitorPauseResponse parses an HTTP response from a ToggleMonitorPauseWithResponse call
func ParseToggleMonitorPauseResponse(rsp *http.Response) (*ToggleMonitorPauseResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ToggleMonitorPauseResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xYUW/bNhD+KwQ3oC+a5a7FUPhpbuwF3pI2SDLsITAMmjpLbChSIU9JjcD/fSAlWZIl",
	"O86WZMDeLJH87u777o4nP1Ku00wrUGjp6JEasJlWFvzDBFYslzg1Rhv3zLVCUOh+siyTgjMUWoXfrFbu",
	"neUJpMz9+tHAio7oD2ENHharNvRol6UZutlsAhqB5UZkDoyO6JjEoMAITsBtJabeG5Q2vHcnBhjCmHOd",
	"K7yEuxysdy0zOgODooiBFesLxVJwz7jOgI6oRSNUTDcBhZQJ2buSMWsftIl6FjcBNXCXCwMRHd20jVSQ",
	"DYD5JijdPddKoIt/j7s8AX67EArB3DO5EGphgWsV+cVUKJHmKR19GAaVR25rDMZHoqJMC4WL3MinfW7t",
	"Dg4Yds63Ret4DVWGtJX0pwjPLeq0FJPrCIjNeUKYJe88Uc5WbuEdDboSpGAti6ELPSaNZ8KWOkeCCRRW",
	"uki7sZe7Kvj59oBefgOOzvQp4FjKUq/P61l0wdZSs6gbfcTwyawvcTqe+LPzXXP2SVsCIbVHG92Gx4xh",
	"6yK342Y5NLLIrSx8Ou9ZB7PYfxo1MlkfX2mTMiy2/PKRdrN2hw4P3LDRcqeNHtTczRQXUdmYdsrJF120",
	"YG1vIobwEwpfrJ2cE0cUvIhcydTgzoszHc/UXuFQ34J6GrjYtoXb2ydeomv1tqkqZ57XmbqJ8E+Yf6KB",
	"7ZEmoKLU3x5dGduM6SkNYRe1I1nD4FJrCUyVezKWW4j6lyWzuPCEPYuBviTbadM7zjU9afJwqJ93Evda",
	"x7GsbqYLB7Y37bypvpg7dez2zf2NDTw3AtdXjvoC5TMwA2acY+Kelv7pt4qf3/+6puU97y341ZqrBDEr",
	"pgb4jmAUkxPNbfeCcPvsKAxjgUm+HHCdhiupH/htGOklmJQpFV5Ox5Pz6SB13PmUO+ZUkXArXQ1EjGOj",
	"IOlKmFQoPeAJUzFT4tfYLTgk2pl0JhXmO0uWjN+Ccp5IwaG8ZYuhhZ7Prp/jYXg2O5l+ufKBufwGk9qv",
	"qysw94LDkUEGFAVKt7uGrV28B2OLEIaD4eC9s6IzUCwTdEQ/DIaDD76tYOKFCVmOSVjOSEUa6SK5XGr5",
	"GXIW0VF7oKNFQoHFzzpav9j02Ts0btrpiyYH/6IxCv88fN/NsZPL6fh6OimE9aPyPvNbrLA1U/v6yNOU",
	"mfU2fksYUfBA2JYHZLF1ReULxtdUwajUsVD76Tzzy69DY+tyOoq+4cvaru7Znu+Hq9npl+mE/HnxUrqc",
	"6Zh4JnuESMuBzRmIoUeF9lzny8KwFBDcmZtH6hSkdzmYNQ2qei9Hn5qO7eD/vm+C6geRIhXYRmHfS5Th",
	"MDiMOX9F/fon3R4hv/7xLxUsbx7Pc/POuZm7AGuBTwEJk5JUYhKhCCMPAhOSsVgoH2VD/62a7vY81Muq",
	"Yeo1e9nOF+V/28uOZbxwvex06ZalHoKbRRY+lr9mk00RgQSELvkT/74mv6/k3O1UF8sWl+6S1yyg3XGt",
	"WyUfu8ROpmfTtyO2CP0wpcHeTlV/7U7elLZXai7Nr/b/WX/pzuyvLdjL96/93x1HNbGeWnszRb23RBuS",
	"K/+p4/97OtzHPLy5r8Spp/lRGErNmUy0xdGn4ach3cw3fwcAAP//zWOyyZgVAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
