package controllers

import (
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/ablancas22/messenger-frontend/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AuthController structures the set of app page views
type AuthController struct {
	manager     *ControllerManager
	authService *services.AuthService
}

// RegisterPage renders Index Page
func (p *AuthController) RegisterPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	rForm := models.InitializeRegistrationForm()
	model := models.IndexModel{Name: "home", Title: "Home", Auth: auth, Form: rForm}
	if auth.Authenticated {
		model.Heading = models.NewHeading("Welcome", "w3-wide text")
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	rModel := models.IndexModel{Name: "register", Title: "Register", Auth: auth, Form: rForm}
	rModel.Heading = models.NewHeading("Register", "w3-wide text")
	p.manager.Viewer.RenderTemplate(w, "templates/register.html", &rModel)
}

// LoginPage renders the Login Page
func (p *AuthController) LoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	lForm := models.InitializeSignInForm()
	lModel := models.LoginModel{Name: "login", Title: "Login", Auth: auth, Form: lForm, Heading: models.NewHeading("Login", "w3-wide text")}
	if auth.Authenticated {
		http.Redirect(w, r, "/", 303)
		return
	}
	p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
}

// LoginHandler controls the login process
func (p *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, cookie := p.manager.authCheck(r)
	lForm := models.InitializeSignInForm()
	model := models.LoginModel{
		Name:    "login",
		Title:   "Login",
		Auth:    auth,
		Form:    lForm,
		Heading: models.NewHeading("Login", "w3-wide text"),
		Status:  true,
		Alert:   models.NewErrorAlert("Invalid Credentials", true),
	}
	if r.Method != http.MethodPost {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	user := &models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	auth, err := p.authService.Authenticate(user)
	if err != nil || auth.Status != http.StatusOK {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	cookie, err = p.manager.SessionManager.NewSession(auth)
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 303)
}

// RegistrationHandler ...
func (p *AuthController) RegistrationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	model := models.IndexModel{
		Name:    "register",
		Title:   "Register",
		Auth:    auth,
		Heading: models.NewHeading("Register", "w3-wide text"),
	}
	if r.Method != http.MethodPost {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	user := &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
	}
	auth, err := p.authService.Register(user)
	if err != nil || auth.Status != http.StatusCreated {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	cookie, err := p.manager.SessionManager.NewSession(auth)
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 303)
}

// LogoutHandler controls the logout process
func (p *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here, Build in Check to ensure logout worked
	auth, cookie := p.manager.authCheck(r)
	if auth.Authenticated {
		auth, _ = p.authService.Invalidate(auth)
		_ = p.manager.SessionManager.DeleteSession(cookie)
	}
	http.Redirect(w, r, "/login", 303)
}
