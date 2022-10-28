package models

// Settings ...
type Settings struct {
	Class        string
	Id           string
	Col1         *Div
	Col3         *Div
	SettingsForm *Form
	Category     string
}

// NewSettings instantiates a default Settings Abstract
func NewSettings(class string, id string, form *Form, col1 *Div, col3 *Div) *Settings {
	return &Settings{
		Class:        class,
		Id:           id,
		Col1:         col1,
		Col3:         col3,
		SettingsForm: form,
		Category:     "default",
	}
}
