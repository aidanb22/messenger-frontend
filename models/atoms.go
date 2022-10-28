package models

/*
Script Types
*/

// Script ...
type Script struct {
	Id       string
	Load     bool
	Category string
}

/*
Heading Types
*/

// Heading ...
type Heading struct {
	Class    string
	Id       string
	Label    string
	Category string
}

// NewHeading ...
func NewHeading(label string, class string) *Heading {
	return &Heading{Class: class, Label: label, Category: "default"}
}

// NewColumnHeading ...
func NewColumnHeading(label string, class string) *Heading {
	return &Heading{Class: class, Label: label, Category: "column"}
}

/*
Alert Types
*/

// Alert ...
type Alert struct {
	task       string
	ClickClose bool
	Category   string
}

// NewAlert constructs and returns a new Link
func NewAlert(task string, cc bool) *Alert {
	return &Alert{
		task:       task,
		ClickClose: cc,
		Category:   "default",
	}
}

// NewSuccessAlert constructs and returns a new Link
func NewSuccessAlert(task string, cc bool) *Alert {
	return &Alert{
		task:       task,
		ClickClose: cc,
		Category:   "success",
	}
}

// NewErrorAlert constructs and returns a new Link
func NewErrorAlert(task string, cc bool) *Alert {
	return &Alert{
		task:       task,
		ClickClose: cc,
		Category:   "error",
	}
}

/*
Link Types
*/

// Link ...
type Link struct {
	Class    string
	Id       string
	Ref      string
	Label    string
	Break    bool
	Category string
}

// NewLink constructs and returns a new Link
func NewLink(class string, id string, ref string, label string, br bool) *Link {
	return &Link{
		Class:    class,
		Id:       id,
		Ref:      ref,
		Label:    label,
		Break:    br,
		Category: "default",
	}
}

/*
Button Types
*/

// Button ...
type Button struct {
	Name     string
	Class    string
	Id       string
	Type     string
	Label    string
	Category string
}

/*
Data Field Types
*/

// SelectOptions ...
type SelectOptions struct {
	Value    string
	Label    string
	Selected bool
}

// GetDataSelectOptions ...
func GetDataSelectOptions[T DataModel](m []T) []*SelectOptions {
	var ops []*SelectOptions
	for _, g := range m {
		ops = append(ops, &SelectOptions{Value: g.GetID(), Label: g.GetLabel()})
	}
	return ops
}

// GetRoleSelectOptions ...
func GetRoleSelectOptions(defVal string) []*SelectOptions {
	if defVal == "" {
		return []*SelectOptions{{Value: "admin", Label: "Admin"}, {Value: "member", Label: "Member"}}
	}
	if defVal == "admin" {
		return []*SelectOptions{{Value: "admin", Label: "Admin", Selected: true}, {Value: "member", Label: "Member"}}
	}
	return []*SelectOptions{{Value: "admin", Label: "Admin"}, {Value: "member", Label: "Member", Selected: true}}
}

// InputField ...
type InputField struct {
	Name     string
	Label    string
	Class    string
	Id       string
	Type     string // acts as the formId in searchBar inputs
	Value    string // acts as the RequestURL in checkbox type
	Multi    bool
	Checked  bool
	Options  []*SelectOptions
	Rows     string
	Cols     string
	Script   *Script
	Category string
}

// NewSearchInput ... (itYPE = formId)
func NewSearchInput(class string, id string, iType string, val string) *InputField {
	return &InputField{
		Name:     "",
		Label:    "",
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Script:   &Script{Category: "searchBar"},
		Category: "search",
	}
}

// NewCheckboxInput ...
func NewCheckboxInput(label string, class string, id string, chk bool, val string) *InputField {
	return &InputField{
		Label:    label,
		Class:    class,
		Id:       id,
		Checked:  chk,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Script:   nil,
		Category: "checkbox",
	}
}

// NewInput ...
func NewInput(name string, label string, class string, id string, iType string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Script:   nil,
		Category: "input",
	}
}

// NewButtonInput ...
func NewButtonInput(name string, label string, class string, id string, iType string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Script:   nil,
		Category: "submit",
	}
}

// NewSelectInput ...
func NewSelectInput(name string, label string, class string, id string, iType string, ops []*SelectOptions, m bool) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    "",
		Multi:    m,
		Options:  ops,
		Rows:     "",
		Cols:     "",
		Script:   nil,
		Category: "select",
	}
}

// NewTextAreaInput ...
func NewTextAreaInput(name string, label string, class string, id string, rows string, cols string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     "",
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     rows,
		Cols:     cols,
		Script:   nil,
		Category: "text",
	}
}

/*
Table Types
*/

// TableData ...
type TableData struct {
	Value    string
	Link     *Link
	Class    string
	Category string
}

// NewTableData ...
func NewTableData(value string, class string) *TableData {
	return &TableData{Value: value, Class: class, Category: "default"}
}

// NewLinkedTableData ...
func NewLinkedTableData(valLink *Link, class string) *TableData {
	return &TableData{Link: valLink, Class: class, Category: "link"}
}

// TableHeader ...
type TableHeader struct {
	Label    string
	Class    string
	Category string
}

// NewTableHeader ...
func NewTableHeader(label string, class string) *TableHeader {
	return &TableHeader{Label: label, Class: class, Category: "default"}
}
