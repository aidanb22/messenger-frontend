package controllers

import (
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/ablancas22/messenger-frontend/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AccountController structures the set of app page views
type AccountController struct {
	manager     *ControllerManager
	userService *services.UserService
	authService *services.AuthService
}

// AccountPage renders the Account Page
func (p *AccountController) AccountPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	up := r.URL.Query().Get("updated")
	model := models.AccountModel{
		Title:    "Account",
		SubRoute: params.ByName("child"),
		Name:     "account",
		Auth:     auth,
	}
	model.Heading = models.NewHeading("My Account", "w3-wide text")
	if model.SubRoute == "settings" {
		user, err := p.userService.Get(auth, &models.User{Id: auth.UserId})
		if err != nil {
			http.Redirect(w, r, "/logout", 303)
			return
		}
		model.User = user
		model.Settings = models.InitializeUserSettings(model.User, false)
		model.Title = "Account Settings"
		model.Name = "Account Settings"
		model.Heading = models.NewHeading("Account Settings", "w3-wide text")
	} else if model.SubRoute == "password" {
		model.PasswordForm = models.InitializePasswordForm()
		model.Title = "Update Password"
		model.Name = "Update Password"
		model.Heading = models.NewHeading("Update Password", "w3-wide text")
	}
	model.BuildRoute()
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if up != "" {
		var alert *models.Alert
		if up == "yes" {
			alert = models.NewSuccessAlert("Account Updated", true)
		} else if up == "no" {
			alert = models.NewErrorAlert("Error Updating Account", true)
		}
		model.Alert = alert
		model.Status = true
	}
	p.manager.Viewer.RenderTemplate(w, "templates/account.html", &model)
}

// AccountSettingsHandler controls the account settings update process
func (p *AccountController) AccountSettingsHandler(w http.ResponseWriter, r *http.Request) {
	upMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	user := &models.User{
		Id:        auth.UserId,
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
	}
	user.GroupId = auth.GroupId
	user, err := p.userService.Update(auth, user)
	if err != nil {
		upMsg = "no"
	}
	http.Redirect(w, r, "/account/settings?updated="+upMsg, 303)
	return
}

// AccountPasswordHandler controls the account settings update process
func (p *AccountController) AccountPasswordHandler(w http.ResponseWriter, r *http.Request) {
	upMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	dto := &models.UpdatePassword{
		NewPassword:     r.FormValue("password"),
		CurrentPassword: r.FormValue("cpassword"),
	}
	_, err := p.authService.UpdatePassword(dto, auth)
	if err != nil {
		upMsg = "no"
	}
	http.Redirect(w, r, "/account/settings?updated="+upMsg, 303)
	return
}
