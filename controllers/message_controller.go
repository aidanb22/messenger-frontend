package controllers

import (
	"encoding/json"
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/ablancas22/messenger-frontend/services"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

//todo: make this messages controller

// taskController structures the set of app page views
type MessageController struct {
	manager        *ControllerManager
	messageService *services.MessageService
	userService    *services.UserService
	groupService   *services.GroupService
}

// tasksPage renders the Variable Page
func (p *MessageController) MessagesPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	created := r.URL.Query().Get("created")
	groupUsers, err := p.groupService.GetGroupUsers(auth, &models.Group{Id: auth.GroupId})
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	userMessages, err := p.messageService.GetMany(auth)
	userMessagesList := models.NewLinkedList(userMessages, "/", true, true, true)
	userMessagesList.Script = &models.Script{Category: "postCheck"}
	createForm := models.InitializePopupCreatetaskForm(groupUsers.Users)
	model := models.MessageModel{
		Title:            "messages",
		SubRoute:         params.ByName("child"),
		Name:             "messages",
		Auth:             auth,
		CreateMessage:    createForm,
		OverviewMessages: userMessagesList,
	}
	model.BuildRoute()
	model.Heading = models.NewHeading("tasks Overview", "w3-wide text")
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if created != "" {
		var alert *models.Alert
		if created == "yes" {
			alert = models.NewSuccessAlert("message Created", true)
		} else if created == "no" {
			alert = models.NewErrorAlert("Error Creating message", true)
		}
		model.Alert = alert
		model.Status = true
	}
	model.Auth = auth
	p.manager.Viewer.RenderTemplate(w, "templates/messages.html", &model)
}

// CreateMessageHandler creates a new user message
func (p *MessageController) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	returnMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	_, err := models.ConvertToDateTime(r.FormValue("due"))
	if err != nil {
		returnMsg = "no"
		http.Redirect(w, r, "/messages?created="+returnMsg, 303)
		return
	}
	messages := &models.Message{
		Id:         r.FormValue("id"),
		SenderID:   r.FormValue("id"),
		ReceiverID: r.FormValue("id"),
		//todo: not sure what to put here, ask
		//Content: r.FormVlue("content"),
		/*
			Id           string    `json:"id,omitempty"`
			SenderID     string    `json:"sender_id,omitempty"`
			ReceiverID   string    `json:"receiver_id,omitempty"`
			Content      string    `json:"content,omitempty"`
			ContentType  string    `json:"contentType,omitempty"`
			Group        bool      `json:"group,omitempty"`
		*/
	}
	/*if !auth.RootAdmin && messages.GroupId != auth.GroupId {
		message.GroupId = auth.GroupId
	}
	*/
	messages, err = p.messageService.Create(auth, messages)
	if err != nil {
		returnMsg = "no"
	}
	http.Redirect(w, r, "/messages?created="+returnMsg, 303)
	return
}

// CompleteMessageHandler updates whether a task is completed or not
func (p *MessageController) CompleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Message
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	if err = r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	if err = json.Unmarshal(body, &t); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	updateId := params.ByName("id")
	t.Id = updateId
	task, err := p.messageService.Update(auth, &t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(task); err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(task); err != nil {
		return
	}
	return
}
