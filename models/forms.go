package models

// InitializeSignInForm for a settings form
func InitializeSignInForm() *Form {
	formMeta := []string{"Sign In", "Auth", "form1", "form1", "POST", "", "default"}
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	fields := []*InputField{emailField, pwField}
	button := &Button{Name: "update", Class: "btn", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil, nil)
}

// InitializeRegistrationForm for a settings form
func InitializeRegistrationForm() *Form {
	formMeta := []string{"Update", "User", "form1", "form1", "POST", "", "default"}
	unField := NewInput("User Name", "User Name", "update", "username", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", "")
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", "")
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	fields := []*InputField{unField, pwField, cpwField, fnField, lnField, emailField}
	button := &Button{Name: "update", Class: "btn", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil, nil)
}

// InitializePasswordForm for updating a user password
func InitializePasswordForm() *Form {
	formMeta := []string{"Update Password", "Password", "form1", "updatePassword", "POST", "", "default"}
	pwField := NewInput("New Password", "New Password", "update", "password", "password", "")
	cpwField := NewInput("Current Password", "Current Password", "update", "cpassword", "password", "")
	fields := []*InputField{pwField, cpwField}
	button := &Button{Name: "update", Class: "btn", Id: "updatePassword", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil, nil)
}

// InitializeUserSettingsForm for a settings form
func InitializeUserSettingsForm(user *User, admin bool) *Form {
	action := "/account/settings"
	if admin {
		action = "/admin/" + user.GetClass(true) + "/" + user.GetID() + "/update"
	}
	formMeta := []string{"Update", "User", "form1", "updateUser", "POST", action, "default"}
	unField := NewInput("User Name", "User Name", "update", "username", "text", user.Username)
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", user.FirstName)
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", user.LastName)
	emailField := NewInput("Email", "Email", "update", "email", "text", user.Email)
	fields := []*InputField{unField, fnField, lnField, emailField}
	if admin {
		pwField := NewInput("Password", "Password", "update", "password", "password", "")
		cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
		roleField := NewSelectInput("User Role", "User Role", "update", "role", "text", GetRoleSelectOptions(user.Role), false)
		fields = append(fields, pwField, cpwField, roleField)
	}
	button := &Button{Name: "update", Class: "btn", Id: "updateUser", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil, nil)
}

// InitializeGroupSettingsForm for a settings form
func InitializeGroupSettingsForm(group *Group) *Form {
	formMeta := []string{"Update", "Group", "form1", "updateGroup", "POST", "/admin/" + group.GetClass(true) + "/" + group.GetID() + "/update", "default"}
	nameField := NewInput("Group Name", "Group Name", "update", "name", "text", group.Name)
	fields := []*InputField{nameField}
	button := &Button{Name: "update", Class: "btn", Id: "updateGroup", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil, nil)
}

// InitializePopupDeleteForm for deletes a data record
func InitializePopupDeleteForm[T DataModel](m T) *Form {
	formId := "delete" + m.GetClass(false) + m.GetID()
	formMeta := []string{"Create", "delete-form", "form-container", formId, "GET", "/admin/" + m.GetClass(true) + "/" + m.GetID() + "/delete", "popup"}
	idField := NewInput("Id", "", "delete", "id", "hidden", m.GetID())
	fields := []*InputField{idField}
	subButton := NewButtonInput("Submit", "", "btn delete", "", "submit", "Delete")
	fields = append(fields, subButton)
	clsButton := &Button{Name: "update", Class: "btn cancel delete", Id: formId, Type: "button", Label: "Cancel", Category: "close-click"}
	opButton := &Button{Name: "delete", Class: "open-button delete", Id: formId, Type: "button", Label: "", Category: "open-click"}
	buttons := []*Button{clsButton}
	return NewForm(formMeta, fields, buttons, opButton, nil)
}

// InitializePopupCreateUserForm for a settings form
func InitializePopupCreateUserForm(availGroups []*Group, setRole bool, master bool) *Form {
	formAction := "/admin/users"
	formMeta := []string{"Create", "User", "form-container", "createUser", "POST"}
	unField := NewInput("User Name", "User Name", "update", "username", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", "")
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", "")
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	fields := []*InputField{unField, pwField, cpwField, fnField, lnField, emailField}
	if len(availGroups) > 1 {
		if master {
			formAction += "?view=table"
		}
		groupField := NewSelectInput("User Group", "User Group", "update", "group_id", "text", GetDataSelectOptions(availGroups), false)
		fields = append(fields, groupField)
	} else if len(availGroups) == 1 {
		formAction = "/admin/groups/" + availGroups[0].Id
		groupField := NewInput("User Group", "", "update", "group_id", "hidden", availGroups[0].Id)
		fields = append(fields, groupField)
	}
	if setRole {
		groupField := NewSelectInput("User Role", "User Role", "update", "role", "text", GetRoleSelectOptions(""), false)
		fields = append(fields, groupField)
	}
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields = append(fields, subButton)
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "createUser", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button create", Id: "createUser", Type: "button", Label: "Create User", Category: "open-click"}
	buttons := []*Button{clsButton}
	formMeta = append(formMeta, formAction)
	formMeta = append(formMeta, "popup")
	return NewForm(formMeta, fields, buttons, opButton, &Script{Load: true, Category: "popupForm"})
}

// InitializePopupCreatetaskForm for a settings form
func InitializePopupCreatetaskForm(availUsers []*User) *Form {
	formAction := "/tasks"
	formMeta := []string{"Create", "task", "form-container", "createtask", "POST"}
	nameField := NewInput("task Name", "task Name", "update", "name", "text", "")
	dueField := NewInput("Due", "Due", "update", "due", "datetime-local", "")
	descField := NewTextAreaInput("Description", "Description", "update", "description", "10", "30", "")
	fields := []*InputField{nameField, dueField, descField}
	if len(availUsers) > 1 {
		userField := NewSelectInput("Assign User", "Assign User", "update", "user_id", "text", GetDataSelectOptions(availUsers), false)
		fields = append(fields, userField)
	} else if len(availUsers) == 1 {
		userField := NewInput("Assign User", "", "update", "user_id", "hidden", availUsers[0].Id)
		fields = append(fields, userField)
	}
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields = append(fields, subButton)
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "createtask", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button create", Id: "createtask", Type: "button", Label: "Create task", Category: "open-click"}
	buttons := []*Button{clsButton}
	formMeta = append(formMeta, formAction)
	formMeta = append(formMeta, "popup")
	return NewForm(formMeta, fields, buttons, opButton, &Script{Load: true, Category: "popupForm"})
}

// InitializePopupCreateGroupForm for a settings form
func InitializePopupCreateGroupForm() *Form {
	formMeta := []string{"Create", "Group", "form-container", "createGroup", "POST", "/admin/groups", "popup"}
	unField := NewInput("Group Name", "Group Name", "update", "name", "text", "")
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields := []*InputField{unField, subButton}
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "createGroup", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button create", Id: "createGroup", Type: "button", Label: "Create Group", Category: "open-click"}
	buttons := []*Button{clsButton}
	return NewForm(formMeta, fields, buttons, opButton, &Script{Load: true, Category: "popupForm"})
}
