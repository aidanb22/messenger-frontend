package server

import (
	"github.com/ablancas22/messenger-frontend/controllers"
	"github.com/julienschmidt/httprouter"
)

// GetRouter returns a new HTTP Router
func GetRouter(p *controllers.ControllerManager, b *controllers.BasicController, au *controllers.AuthController,
	ac *controllers.AccountController, ad *controllers.AdminController, t *controllers.MessageController) *httprouter.Router {

	// mux handler
	router := httprouter.New()

	// Index route
	router.GET("/", b.IndexPage)
	router.GET("/register", au.RegisterPage)
	// Register new Account/User at Index Page Form
	router.POST("/register", au.RegistrationHandler)

	// Login Route
	router.GET("/login", au.LoginPage)
	router.POST("/login", au.LoginHandler)

	// Logout Route
	router.GET("/logout", au.LogoutHandler)

	// About Route
	router.GET("/about", b.AboutPage)
	router.GET("/about/:child", b.AboutPage)

	// Account Route
	router.Handler("GET", "/account", p.Protected(ac.AccountPage))
	router.Handler("GET", "/account/:child", p.Protected(ac.AccountPage))

	// Account Settings Route
	router.Handler("POST", "/account/settings", p.Protected(ac.AccountSettingsHandler))
	router.Handler("POST", "/account/password", p.Protected(ac.AccountPasswordHandler))

	// Admin Page Routes
	router.Handler("GET", "/admin", p.Protected(ad.AdminPage))
	router.Handler("GET", "/admin/users-table", p.Protected(ad.MasterAdminUserPage))
	router.Handler("GET", "/admin/users/:id", p.Protected(ad.AdminUserPage))

	// Admin Group Handler Routes
	router.Handler("GET", "/admin/groups", p.Protected(ad.AdminGroupsPage))
	router.Handler("GET", "/admin/groups/:id", p.Protected(ad.AdminGroupPage))
	router.Handler("POST", "/admin/groups", p.Protected(ad.AdminCreateGroupHandler))
	router.Handler("POST", "/admin/groups/:id/update", p.Protected(ad.AdminUpdateGroupHandler))
	router.Handler("GET", "/admin/groups/:id/delete", p.Protected(ad.AdminDeleteGroupHandler))

	// Admin User Handler Routes
	router.Handler("POST", "/admin/groups/:id", p.Protected(ad.AdminCreateUserHandler))
	router.Handler("POST", "/admin/users", p.Protected(ad.AdminCreateUserHandler))
	router.Handler("POST", "/admin/users/:id/update", p.Protected(ad.AdminUpdateUserHandler))
	router.Handler("GET", "/admin/users/:id/delete", p.Protected(ad.AdminDeleteUserHandler))

	// task Page Routes
	router.Handler("GET", "/tasks", p.Protected(t.MessagesPage))
	//router.Handler("GET", "/tasks/:id", p.Protected(b.taskPage))

	// task Handler Routes
	router.Handler("POST", "/tasks", p.Protected(t.CreateMessageHandler))
	router.Handler("POST", "/tasks/:id/check", p.Protected(t.CompleteMessageHandler))

	// Example route that encounters an error
	router.GET("/broken/handler", b.BrokenPage)

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", p.Viewer.Statics)

	return router
}
