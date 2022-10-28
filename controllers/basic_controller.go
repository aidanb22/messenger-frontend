package controllers

import (
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// BasicController structures the set of app page views
type BasicController struct {
	manager *ControllerManager
}

// BrokenPage renders Broken Page - for missing/error routes
func (p *BasicController) BrokenPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p.manager.Viewer.RenderTemplate(w, "templates/missing.html", nil)
}

// IndexPage renders Index Page
func (p *BasicController) IndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, cookie := p.manager.authCheck(r)
	model := models.IndexModel{Name: "register", Title: "Register", Auth: auth}
	if !auth.Authenticated {
		model.Form = models.InitializeRegistrationForm()
		model.Heading = models.NewHeading("Create an account", "w3-wide text")
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	auth, err := p.manager.SessionManager.GetSession(cookie)
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	model.Name = "Home"
	model.Title = "Home"
	model.Auth = auth
	model.Heading = models.NewHeading("Welcome", "w3-wide text")
	p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
}

// AboutPage renders the About Page
func (p *BasicController) AboutPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth, _ := p.manager.authCheck(r)
	model := models.AboutModel{Title: "About", SubRoute: ps.ByName("child"), Name: "about", Auth: auth}
	model.BuildRoute()
	model.Heading = models.NewHeading("About", "w3-wide text")
	if !auth.Authenticated {
		p.manager.Viewer.RenderTemplate(w, "templates/about.html", &model)
		return
	}
	model.Auth = auth
	p.manager.Viewer.RenderTemplate(w, "templates/about.html", &model)
}

// VariablePage renders the Variable Page
func (p *BasicController) VariablePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth, _ := p.manager.authCheck(r)
	model := models.VariableModel{Title: "Variable", SubRoute: ps.ByName("child"), Name: "variable", Auth: auth}
	model.BuildRoute()
	model.Heading = models.NewHeading("Variable", "w3-wide text")
	if !auth.Authenticated {
		p.manager.Viewer.RenderTemplate(w, "templates/variable.html", &model)
		return
	}
	model.Auth = auth
	p.manager.Viewer.RenderTemplate(w, "templates/variable.html", &model)
}
