package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/JECSand/fetch"
	"github.com/ablancas22/messenger-frontend/models"
	"io"
	"net/http"
	"os"
)

/*
================ APIRequest ==================
*/

// Request managers an async request to the API to return data model(s)
type Request struct {
	url       string
	authToken string
}

// getAuthHeaders builds a fetch usable slice of request headers
func (dr *Request) getAuthHeaders() [][]string {
	reqHeaders := fetch.JSONDefaultHeaders()
	if dr.authToken != "" {
		return fetch.AppendHeaders(reqHeaders, []string{"Auth-Token", dr.authToken})
	}
	return reqHeaders
}

// Get returns one or more dataModel
func (dr *Request) Get() (*fetch.Fetch, error) {
	f, err := fetch.NewFetch(dr.url, "GET", dr.getAuthHeaders(), nil)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Post a new dataModel
func (dr *Request) Post(bodyContents []byte) (*fetch.Fetch, error) {
	body := bytes.NewBuffer(bodyContents)
	f, err := fetch.NewFetch(dr.url, "POST", dr.getAuthHeaders(), body)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Patch an existing data model
func (dr *Request) Patch(bodyContents []byte) (*fetch.Fetch, error) {
	body := bytes.NewBuffer(bodyContents)
	f, err := fetch.NewFetch(dr.url, "PATCH", dr.getAuthHeaders(), body)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Delete a dataModel
func (dr *Request) Delete() (*fetch.Fetch, error) {
	f, err := fetch.NewFetch(dr.url, "DELETE", dr.getAuthHeaders(), nil)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// NewRequest initializes and returns a new Request struct
func NewRequest(url string, token string) *Request {
	return &Request{url: url, authToken: token}
}

/*
================ API Request ==================
*/

// APIRequest is a Generic type struct for organizing dataModel methods
type APIRequest[T models.DTOModel] struct {
	url  string
	auth *models.Auth
}

// LoadModel loads returned json data into a dataModel
func (api *APIRequest[T]) LoadModel(resp *http.Response) (T, error) {
	var m T
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
		return m, errors.New("response status error")
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return m, err
	}
	if err = resp.Body.Close(); err != nil {
		return m, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// Get a dataModel
func (api *APIRequest[T]) Get() (T, error) {
	var m T
	newReq := NewRequest(api.url, api.auth.AuthToken)
	f, err := newReq.Get()
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.LoadModel(f.Res)
}

// GetAsync a dataModel
func (api *APIRequest[T]) GetAsync() (*fetch.Fetch, error) {
	return NewRequest(api.url, api.auth.AuthToken).Get()
}

// Create a dataModel
func (api *APIRequest[T]) Create(data T) (T, error) {
	var m T
	newReq := NewRequest(api.url, api.auth.AuthToken)
	f, err := newReq.Post(data.GetJSON())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.LoadModel(f.Res)
}

// Update a dataModel
func (api *APIRequest[T]) Update(data T) (T, error) {
	var m T
	newReq := NewRequest(api.url, api.auth.AuthToken)
	f, err := newReq.Patch(data.GetJSON())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.LoadModel(f.Res)
}

// Delete a dataModel
func (api *APIRequest[T]) Delete() (T, error) {
	var m T
	newReq := NewRequest(api.url, api.auth.AuthToken)
	f, err := newReq.Delete()
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.LoadModel(f.Res)
}

/*
================ User Service ==================
*/

// UserService is a Generic type struct for organizing dataModel methods
type UserService struct {
	host     string
	endpoint string
}

// NewUsersRequest ... where filter = "/" + filter.GetID() or just ""
func (s *UserService) NewUsersRequest(filter string, auth *models.Auth) *APIRequest[*models.Users] {
	return &APIRequest[*models.Users]{url: s.host + s.endpoint + filter, auth: auth}
}

// NewUserRequest ... filter string,
func (s *UserService) NewUserRequest(filter string, auth *models.Auth) *APIRequest[*models.User] {
	return &APIRequest[*models.User]{url: s.host + s.endpoint + filter, auth: auth}
}

// GetMany returns a slice of dataModels
func (s *UserService) GetMany(auth *models.Auth) ([]*models.User, error) {
	req := &APIRequest[*models.Users]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Get()
	if err != nil {
		return []*models.User{}, err
	}
	return f.Items, nil
}

// Get a dataModel
func (s *UserService) Get(auth *models.Auth, filter *models.User) (*models.User, error) {
	req := &APIRequest[*models.User]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Get()
	if err != nil {
		return &models.User{}, err
	}
	return f, nil
}

// Create a dataModel
func (s *UserService) Create(auth *models.Auth, data *models.User) (*models.User, error) {
	req := &APIRequest[*models.User]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Create(data)
	if err != nil {
		return &models.User{}, err
	}
	return f, nil
}

// Update a dataModel
func (s *UserService) Update(auth *models.Auth, data *models.User) (*models.User, error) {
	req := &APIRequest[*models.User]{url: s.host + s.endpoint + "/" + data.GetID(), auth: auth}
	f, err := req.Update(data)
	if err != nil {
		return &models.User{}, err
	}
	return f, nil
}

// Delete a dataModel
func (s *UserService) Delete(auth *models.Auth, filter *models.User) (*models.User, error) {
	req := &APIRequest[*models.User]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Delete()
	if err != nil {
		return &models.User{}, err
	}
	return f, nil
}

// NewUserService initializes and returns a new APIHandler
func NewUserService() *UserService {
	return &UserService{
		host:     os.Getenv("API_HOST"),
		endpoint: "/users",
	}
}

/*
================ Group Service ==================
*/

// GroupService is a Generic type struct for organizing dataModel methods
type GroupService struct {
	host     string
	endpoint string
}

// NewGroupsRequest ...
func (s *GroupService) NewGroupsRequest(filter string, auth *models.Auth) *APIRequest[*models.Groups] {
	return &APIRequest[*models.Groups]{url: s.host + s.endpoint + filter, auth: auth}
}

// NewGroupRequest ...
func (s *GroupService) NewGroupRequest(filter string, auth *models.Auth) *APIRequest[*models.Group] {
	return &APIRequest[*models.Group]{url: s.host + s.endpoint + filter, auth: auth}
}

// GetMany returns a slice of dataModels
func (s *GroupService) GetMany(auth *models.Auth) ([]*models.Group, error) {
	req := &APIRequest[*models.Groups]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Get()
	if err != nil {
		return []*models.Group{}, err
	}
	return f.Items, nil
}

// Get a dataModel
func (s *GroupService) Get(auth *models.Auth, filter *models.Group) (*models.Group, error) {
	req := &APIRequest[*models.Group]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Get()
	if err != nil {
		return &models.Group{}, err
	}
	return f, nil
}

// Create a dataModel
func (s *GroupService) Create(auth *models.Auth, data *models.Group) (*models.Group, error) {
	req := &APIRequest[*models.Group]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Create(data)
	if err != nil {
		return &models.Group{}, err
	}
	return f, nil
}

// Update a dataModel
func (s *GroupService) Update(auth *models.Auth, data *models.Group) (*models.Group, error) {
	req := &APIRequest[*models.Group]{url: s.host + s.endpoint + "/" + data.GetID(), auth: auth}
	f, err := req.Update(data)
	if err != nil {
		return &models.Group{}, err
	}
	return f, nil
}

// Delete a dataModel
func (s *GroupService) Delete(auth *models.Auth, filter *models.Group) (*models.Group, error) {
	req := &APIRequest[*models.Group]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Delete()
	if err != nil {
		return &models.Group{}, err
	}
	return f, nil
}

// GetGroupUsers returns a models.GroupUsersDTO struct dataModel
func (s *GroupService) GetGroupUsers(auth *models.Auth, filter *models.Group) (*models.GroupUsersDTO, error) {
	req := &APIRequest[*models.GroupUsersDTO]{url: s.host + s.endpoint + "/" + filter.GetID() + "/users", auth: auth}
	f, err := req.Get()
	if err != nil {
		return &models.GroupUsersDTO{}, err
	}
	return f, nil
}

// NewGroupService initializes and returns a new APIHandler
func NewGroupService() *GroupService {
	return &GroupService{
		host:     os.Getenv("API_HOST"),
		endpoint: "/groups",
	}
}

/*
================ Group Service ==================
*/

// taskService is a Generic type struct for organizing dataModel methods
type MessageService struct {
	host     string
	endpoint string
}

// NewtasksRequest ...
func (s *MessageService) NewtasksRequest(filter string, auth *models.Auth) *APIRequest[*models.Messages] {
	return &APIRequest[*models.Messages]{url: s.host + s.endpoint + filter, auth: auth}
}

// NewtaskRequest ...
func (s *MessageService) NewtaskRequest(filter string, auth *models.Auth) *APIRequest[*models.Message] {
	return &APIRequest[*models.Message]{url: s.host + s.endpoint + filter, auth: auth}
}

// GetMany returns a slice of dataModels
func (s *MessageService) GetMany(auth *models.Auth) ([]*models.Message, error) {
	req := &APIRequest[*models.Messages]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Get()
	if err != nil {
		return []*models.Message{}, err
	}
	return f.Items, nil
}

// Get a dataModel
func (s *MessageService) Get(auth *models.Auth, filter *models.Message) (*models.Message, error) {
	req := &APIRequest[*models.Message]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Get()
	if err != nil {
		return &models.Message{}, err
	}
	return f, nil
}

// Create a dataModel
func (s *MessageService) Create(auth *models.Auth, data *models.Message) (*models.Message, error) {
	req := &APIRequest[*models.Message]{url: s.host + s.endpoint, auth: auth}
	f, err := req.Create(data)
	if err != nil {
		return &models.Message{}, err
	}
	return f, nil
}

// Update a dataModel
func (s *MessageService) Update(auth *models.Auth, data *models.Message) (*models.Message, error) {
	req := &APIRequest[*models.Message]{url: s.host + s.endpoint + "/" + data.GetID(), auth: auth}
	f, err := req.Update(data)
	if err != nil {
		return &models.Message{}, err
	}
	return f, nil
}

// Delete a dataModel
func (s *MessageService) Delete(auth *models.Auth, filter *models.Message) (*models.Message, error) {
	req := &APIRequest[*models.Message]{url: s.host + s.endpoint + "/" + filter.GetID(), auth: auth}
	f, err := req.Delete()
	if err != nil {
		return &models.Message{}, err
	}
	return f, nil
}

// NewtaskService initializes and returns a new APIHandler
func NewtaskService() *MessageService {
	return &MessageService{
		host:     os.Getenv("API_HOST"),
		endpoint: "/messages",
	}
}
