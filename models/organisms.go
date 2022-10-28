package models

/*
Forms
*/

// Form structures a generic form
type Form struct {
	Name     string        `json:"name,omitempty"`
	Type     string        `json:"type,omitempty"`
	Class    string        `json:"class,omitempty"`
	Id       string        `json:"id,omitempty"`
	Inputs   []*InputField `json:"inputs,omitempty"`
	Buttons  []*Button     `json:"buttons,omitempty"`
	Popup    *Button       `json:"popup,omitempty"`
	Script   *Script       `json:"script,omitempty"`
	Method   string        `json:"method,omitempty"`
	Action   string        `json:"action,omitempty"`
	Category string        `json:"category,omitempty"`
}

// NewForm for a new form
func NewForm(formMeta []string, fields []*InputField, buttons []*Button, popup *Button, script *Script) *Form {
	newForm := Form{}
	// Field Vector String Array, this is order
	// Name, Class, Id, Type, Label, DefaultVal
	newForm.Name = formMeta[0]
	newForm.Type = formMeta[1]
	newForm.Class = formMeta[2]
	newForm.Id = formMeta[3]
	newForm.Method = formMeta[4]
	newForm.Action = formMeta[5]
	newForm.Inputs = fields
	newForm.Buttons = buttons
	newForm.Category = formMeta[6]
	if popup != nil {
		newForm.Popup = popup
	}
	newForm.Script = script
	return &newForm
}

/*
Lists
*/

// List structures a generic form
type List struct {
	Class     string      `json:"class,omitempty"`
	Id        string      `json:"id,omitempty"`
	Items     []*ListItem `json:"inputs,omitempty"`
	SearchBar *InputField `json:"search_bar,omitempty"`
	Script    *Script     `json:"script,omitempty"`
	Category  string      `json:"category,omitempty"`
}

// NewUnorderedList for a new ul
func NewUnorderedList(class string, id string, items []*ListItem, filter *InputField) *List {
	return &List{
		Class:     class,
		Id:        id,
		Items:     items,
		SearchBar: filter,
		Category:  "unordered",
	}
}

// NewOrderedList for a new ul
func NewOrderedList(class string, id string, items []*ListItem) *List {
	return &List{
		Class:    class,
		Id:       id,
		Items:    items,
		Category: "ordered",
	}
}
