package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Auth models the authentication structure of an app user
type Auth struct {
	Username      string
	UserId        string
	GroupId       string
	Role          string
	AuthToken     string
	APIKey        string
	Authenticated bool
	RootAdmin     bool
	LastLogin     string
	LoginIP       string
	Status        int
}

// GetAuthString converts the structures attributes into a string (todo - look into using json marshall/unmarshall)
func (a *Auth) GetAuthString() string {
	authStr := "false"
	rootStr := "false"
	if a.Authenticated {
		authStr = "true"
	}
	if a.RootAdmin {
		rootStr = "true"
	}
	return a.Username + "||" + a.UserId + "||" + a.GroupId + "||" + a.Role + "||" + authStr + "||" + rootStr + "||" + a.AuthToken + "||" + a.LastLogin + "||" + a.LoginIP
}

// LoadAuthString loads a stringed Auth struct
func (a *Auth) LoadAuthString(authString string) {
	authBool := false
	rootBool := false
	sString := strings.Split(authString, "||")
	authStr := sString[4]
	rootStr := sString[5]
	if authStr == "true" {
		authBool = true
	}
	if rootStr == "true" {
		rootBool = true
	}
	a.Username = sString[0]
	a.UserId = sString[1]
	a.GroupId = sString[2]
	a.Role = sString[3]
	a.Authenticated = authBool
	a.RootAdmin = rootBool
	a.AuthToken = sString[6]
	a.LastLogin = sString[7]
	a.LoginIP = sString[8]
}

// Delete all data in Auth struct
func (a *Auth) Delete() {
	a.Username = ""
	a.UserId = ""
	a.GroupId = ""
	a.Role = ""
	a.Authenticated = false
	a.RootAdmin = false
	a.AuthToken = ""
	a.LastLogin = ""
	a.LoginIP = ""
}

// Load Auth with data from a http.Response
func (a *Auth) Load(resp *http.Response) error {
	currentTime := time.Now().UTC()
	authToken := resp.Header["Auth-Token"]
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return err
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}
	a.Status = resp.StatusCode
	if a.Status == http.StatusOK || a.Status == http.StatusCreated {
		a.Username = user.Username
		a.UserId = user.Id
		a.GroupId = user.GroupId
		a.Role = user.Role
		a.AuthToken = authToken[0]
		a.LastLogin = currentTime.String()
		a.Authenticated = true
		a.RootAdmin = user.RootAdmin
		a.LoginIP = "Unknown"
		return nil
	}
	codeStr := strconv.Itoa(a.Status)
	return errors.New("auth request status returned with error status of " + codeStr)
}
