package server

import (
	"flag"
	"fmt"
	"github.com/ablancas22/messenger-frontend/controllers"
	"github.com/ablancas22/messenger-frontend/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Server is a struct that stores the API Apps high level attributes such as the router, config, and services
type Server struct {
	manager *controllers.ControllerManager
	Router  *httprouter.Router
}

// NewServer is a function used to initialize a new Server struct
func NewServer(manager *controllers.ControllerManager, auth *services.AuthService, u *services.UserService, g *services.GroupService, t *services.taskService) *Server {
	s := Server{manager: manager}
	basicController := s.manager.NewBasicController()
	authController := s.manager.NewAuthController(auth)
	accountController := s.manager.NewAccountController(u, auth)
	adminController := s.manager.NewAdminController(u, g)
	taskController := s.manager.NewMessageController(u, g, t)
	s.Router = GetRouter(manager, basicController, authController, accountController, adminController, taskController)
	return &s
}

// Start starts the initialized server
func (s *Server) Start() {
	port := ":" + os.Getenv("PORT")
	listen := flag.String("listen ", port, "Interface and port to listen on")
	flag.Parse()
	fmt.Println("Listening on ", *listen)
	go func() {
		log.Fatal(http.ListenAndServe(*listen, s.Router))
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
}
