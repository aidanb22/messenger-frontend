package controllers

import (
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/ablancas22/messenger-frontend/services"
	"github.com/ablancas22/messenger-frontend/views"
	"net/http"
	"strings"
)

// splitPathURL
func splitPathURL(r *http.Request) []string {
	reqUrl := r.URL.Path
	var splitURL []string
	splitURL = strings.Split(reqUrl, "/")
	splitURL = splitURL[1:]
	return splitURL
}

// ControllerManager structures the set of app page views
type ControllerManager struct {
	Viewer         *views.Viewer
	SessionManager *services.SessionService
}

// NewManager initializes a new controller manager
func NewManager(viewer *views.Viewer, sm *services.SessionService) *ControllerManager {
	return &ControllerManager{
		Viewer:         viewer,
		SessionManager: sm,
	}
}

// authCheck
func (p *ControllerManager) authCheck(r *http.Request) (*models.Auth, *http.Cookie) {
	cookie, err := r.Cookie("SessionID")
	auth := &models.Auth{Authenticated: false}
	if err != nil {
		return auth, cookie
	}
	authenticated := p.SessionManager.IsLoggedIn(r)
	if err != nil || !authenticated {
		return auth, cookie
	}
	auth, _ = p.SessionManager.GetSession(cookie)
	return auth, cookie
}

// Protected ensures that a user is logged in to view page
func (p *ControllerManager) Protected(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, _ := p.authCheck(r)
		if !auth.Authenticated {
			http.Redirect(w, r, "/login", 303)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

// NewBasicController initialized a BasicViews struct for rendering Basic Views
func (p *ControllerManager) NewBasicController() *BasicController {
	return &BasicController{
		manager: p,
	}
}

// NewAuthController initialized a AuthViews struct for rendering Auth Views
func (p *ControllerManager) NewAuthController(as *services.AuthService) *AuthController {
	return &AuthController{
		manager:     p,
		authService: as,
	}
}

// NewAccountController initialized a BasicViews struct for rendering Account Views
func (p *ControllerManager) NewAccountController(uService *services.UserService, as *services.AuthService) *AccountController {
	return &AccountController{
		manager:     p,
		userService: uService,
		authService: as,
	}
}

// NewAdminController initialized a BasicViews struct for rendering Admin Views
func (p *ControllerManager) NewAdminController(uService *services.UserService, gService *services.GroupService) *AdminController {
	return &AdminController{
		manager:      p,
		userService:  uService,
		groupService: gService,
	}
}

// NewtaskController initialized a BasicViews struct for rendering Admin Views
func (p *ControllerManager) NewMessageController(uService *services.UserService, gService *services.GroupService, tService *services.MessageService) *MessageController {
	return &MessageController{
		manager:        p,
		messageService: tService,
		userService:    uService,
		groupService:   gService,
	}
}
