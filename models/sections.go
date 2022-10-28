package models

// InitializeUserSettings ...
func InitializeUserSettings(user *User, admin bool) *Settings {
	SettingsForm := InitializeUserSettingsForm(user, admin)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	var links []*Link
	infoLink := NewLink("active", "", "/account/settings", "Change User Info", true)
	links = append(links, infoLink)
	if !admin {
		pwLink := NewLink("", "", "/account/password", "Change Password", true)
		links = append(links, pwLink)
	}
	optionsCol := NewLinkDiv("columnOne", "", "", NewColumnHeading("Options", ""), links)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	unLink := NewLink("active", "", "", user.Username, true)
	rLink := NewLink("active", "", "", user.Role, true)
	infoCol := NewLinkDiv("columnTwo", "", "", NewColumnHeading("Current Info", ""), []*Link{unLink, rLink})
	return NewSettings("", "", SettingsForm, optionsCol, infoCol)
}

// InitializeGroupSettings ...
func InitializeGroupSettings(group *Group, users []*User) *Settings {
	SettingsForm := InitializeGroupSettingsForm(group)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	uList := NewLinkedList(users, "/admin/", true, true, false)
	usersCol := NewListDiv("columnOne", "", "", NewColumnHeading("Group Users", ""), uList)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	nameLink := NewLink("active", "", "", group.Name, true)
	infoCol := NewLinkDiv("columnTwo", "", "", NewColumnHeading("Group Info", ""), []*Link{nameLink})
	return NewSettings("", "", SettingsForm, usersCol, infoCol)
}
