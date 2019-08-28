//Package kazoo implements 2600hz Kazoo project API
//and might be used for variety of operations.
//I'm aimed to support V2 version of the API, though
//it might work for V1 API as well
package kazooapi

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	apiURL = "http://localhost:8000/v2"
	token  = ""
	ver    = "1.0.1"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

//RequestEnvelope is the structure which must be presented
//for each request containing a body: POST,PUT etc
type RequestEnvelope struct {
	Data      interface{} `json:"data"`
	AuthToken string      `json:"auth_token,omitempty"` //optional
	Verb      string      `json:"verb,omitempty"`       //optional
}

//ResponseEnvelope is the main structure which must be presented
//for each request contains body: POST,PUT
type ResponseEnvelope struct {
	//Data      interface{} `json:"data"`
	AuthToken string `json:"auth_token,omitempty"`
	Status    string `json:"status,omitempty"`     //one of "success", "error" or "fatal"
	Message   string `json:"message,omitempty"`    //optional message that's should clarify
	Error     string    `json:"error,omitempty"`      //error code
	RequestID string `json:"request_id,omitempty"` //for debugging purposes
	PageSize  int    `json:"page_size,omitempty"`
	Revision  string `json:"revision"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	Node      string `json:"node"`
}

//ErrorResponseEnvelope represents Error data recieved from Kazoo (in case if response code >= 300)
type ErrorResponseEnvelope struct {
	Data interface{} `json:"data"`
	ResponseEnvelope
}

//AuthResponse represents a Kazoo authentication record
type AuthResponse struct {
	PageSize  int                    `json:"page_size"`
	Data      map[string]interface{} `json:"data"`
	Revision  string                 `json:"revision"`
	Timestamp string                 `json:"timestamp"`
	Version   string                 `json:"version"`
	Node      string                 `json:"node"`
	RequestID string                 `json:"request_id"`
	Status    string                 `json:"status"`
	AuthToken string                 `json:"auth_token"`
}

//Paginator is structure which contains data and pagination features
type Paginator struct {
	PageSize int64
	HasNext  bool
	HasPrev  bool
}

/*func (p *Paginator) Next() {

}*/

// APIClient manages communication with a Kazoo API server
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AccountsAPI     *AccountsAPIService
	ChannelsAPI     *ChannelsAPIService
	RecordingsAPI   *RecordingsAPIService
	PhoneNumbersAPI *PhoneNumbersAPIService
	UsersAPI        *UsersAPIService
	DevicesAPI      *DevicesAPIService
	CallflowsAPI    *CallflowsAPIService
	//SupAPI          *SupApiService
	StorageAPI *StorageAPIService
	LimitsAPI  *LimitsAPIService
}

type service struct {
	client *APIClient
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := gregorianToUnixString(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(*v)
	return nil
}

// NewAPIClient creates a new API client.
//Requires an API key or a set of username/password credentials
//Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) (api *APIClient, err error) {

	if cfg.APIKey == "" {
		//return nil, reportError("You have to specify an API key or username/password/realm")
		if cfg.BasicAuth.Username == "" || cfg.BasicAuth.Password == "" || cfg.BasicAuth.Realm == "" {
			return nil, reportError("you have to specify APIKey or username/password/realm")
		}
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
		cfg.HTTPClient.Timeout = time.Second * 5
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccountsAPI = (*AccountsAPIService)(&c.common)
	c.ChannelsAPI = (*ChannelsAPIService)(&c.common)
	c.RecordingsAPI = (*RecordingsAPIService)(&c.common)
	c.PhoneNumbersAPI = (*PhoneNumbersAPIService)(&c.common)
	c.UsersAPI = (*UsersAPIService)(&c.common)
	c.DevicesAPI = (*DevicesAPIService)(&c.common)
	c.CallflowsAPI = (*CallflowsAPIService)(&c.common)
	c.StorageAPI = (*StorageAPIService)(&c.common)
	c.LimitsAPI = (*LimitsAPIService)(&c.common)

	return c, nil
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func readBody(resp *http.Response, v interface{}) error {
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		resp.Body.Close()
		return err
	}
	return resp.Body.Close()
}

func path(segments ...string) string {
	r := ""
	for _, seg := range segments {
		r += "/"
		r += url.QueryEscape(seg)
	}
	return r
}

// callAPI do the request.
func (c *APIClient) callAPI(ctx context.Context, request *http.Request) (resp *http.Response, err error) {

	if !checkAuthState() {
		if err := c.Authenticate(ctx); err != nil {
			return nil, err
		}
		request.Header.Add("X-Auth-Token", token)
		resp, err = c.cfg.HTTPClient.Do(request)
	} else {
		request.Header.Add("X-Auth-Token", token)
		resp, err = c.cfg.HTTPClient.Do(request)
	}

	if resp != nil {
		switch resp.StatusCode {
		case 401:
			go c.Authenticate(ctx)
			request.Header.Add("X-Auth-Token", token)
			resp, err = c.cfg.HTTPClient.Do(request)
			if err != nil {
				return nil, err
			}
		case 0:
			return nil, errors.New("Have not recieved a response from the server")
		}
	}

	return resp, nil
}

//ChangeBasePath enables switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

//Request is the structure which represents all necessary params for makeing
//a new request
type Request struct {
	CTX          context.Context
	Path         string
	Method       string
	PostBody     interface{}
	HeaderParams map[string]string
	QueryParams  url.Values
	FileName     string
	FileBytes    []byte
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(req *Request) (httpRequest *http.Request, err error) {

	var body *bytes.Buffer

	//Initialize empty map of headers
	req.HeaderParams = make(map[string]string)

	// Detect postBody type and post.
	if req.PostBody != nil {
		contentType := req.HeaderParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(req.PostBody)
			req.HeaderParams["Content-Type"] = contentType
		}

		body, err = setBody(req.PostBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	url, err := url.Parse(req.Path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range req.QueryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		httpRequest, err = http.NewRequest(req.Method, url.String(), body)
	} else {
		httpRequest, err = http.NewRequest(req.Method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(req.HeaderParams) > 0 {
		headers := http.Header{}
		for h, v := range req.HeaderParams {
			headers.Set(h, v)
		}
		httpRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		httpRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	httpRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	ctx := req.CTX

	if ctx != nil {
		// add context to the request
		httpRequest = httpRequest.WithContext(ctx)

		// Walk through any authentication.

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			httpRequest.SetBasicAuth(auth.Username, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(string); ok {
			httpRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		httpRequest.Header.Add(header, value)
	}

	return httpRequest, nil
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

//If we get response with error code (i.e. >=300) we can easily report error
func prepareError(resp *http.Response) error {
	var errData	ErrorResponseEnvelope

	err := readBody(resp, &errData)
	if err != nil {
		return reportError("Can't decode error response: %v", err)
	}

	return reportError("Code: %v, Message: %s, Problem: %#v", resp.Status, errData.Message, errData.Data)
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

//Authenticate provides an authentication on a
//Kazoo API server for both api_key and password
//authentication methods
func (c *APIClient) Authenticate(ctx context.Context) error {
	req := Request{
		Method: "PUT",
		CTX:    ctx,
	}

	type authData struct {
		Credentials string `json:"credentials,omitempty"`
		APIKey      string `json:"api_key,omitempty"`
	}

	//Authenticate requests using username/password/realm chain
	//it slightly different from the APIAuth

	ad := authData{}

	if c.cfg.APIKey != "" {
		req.Path = c.cfg.BasePath + "/api_auth"
		ad.APIKey = c.cfg.APIKey
	} else {
		if c.cfg.BasicAuth.Username != "" || c.cfg.BasicAuth.Password != "" || c.cfg.BasicAuth.Realm != "" {
			req.Path = c.cfg.BasePath + "/user_auth"

			hasher := md5.New()
			hash := hasher.Sum([]byte(c.cfg.BasicAuth.Username + ":" + c.cfg.BasicAuth.Password))
			ad.Credentials = hex.EncodeToString(hash)
		} else {
			return reportError("")
		}
	}

	auth := RequestEnvelope{
		Data: ad,
	}

	authBody, jsonErr := json.Marshal(auth)
	if jsonErr != nil {
		return jsonErr
	}

	postBody, bodyErr := setBody(authBody, "json")
	if bodyErr != nil {
		return bodyErr
	}

	// create path and map variables
	req.PostBody = postBody
	req.HeaderParams = make(map[string]string)
	req.QueryParams = url.Values{}

	// body params
	//localVarPostBody = &body
	r, err := c.common.client.prepareRequest(&req)

	if err != nil {
		return err
	}

	authResponse, err := c.cfg.HTTPClient.Do(r)
	//authResponse, err := c.callAPI(ctx, r)
	if err != nil || authResponse == nil {
		return err
	}
	defer authResponse.Body.Close()

	switch authResponse.StatusCode {
	case 201:
		//Succesfull authentication
		authdata := AuthResponse{}
		readBody(authResponse, &authdata)

		token = authdata.AuthToken
		go authTokenExpire()

		return nil

	default:
		//bodyBytes, _ := ioutil.ReadAll(authResponse.Body)
		return reportError("wrong response code during authentication: %d", authResponse.StatusCode)

		//return reportError("Wrong authorization response from the server")
	}

	//return nil
}

func authTokenExpire() {
	timer := time.NewTimer(60 * time.Minute)
	<-timer.C
	token = ""
}

func checkAuthState() bool {
	if len(token) == 0 {
		return false
	}
	return true
}

func gregorianToUnixString(greg string) (unix *time.Time, err error) {
	var gregSecondsSinceUnix int64 = 62167219200

	gDate, err := strconv.ParseInt(greg, 10, 64)
	if err != nil {
		return nil, errors.New("can't parse given string to int64")
	}

	//If we want to know how many seconds passed, then we have to minus amount of seconds passed till the Unix epoch
	unixTimestamp := gDate - gregSecondsSinceUnix

	ts := time.Unix(unixTimestamp, 0)

	return &ts, nil
}
