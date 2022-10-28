package models

// Table ...
type Table struct {
	Class     string
	Id        string
	TableHead *TableRow
	TableBody *TableBody
	Script    *Script
	Category  string
}

// NewTable ...
func NewTable(class string, id string, thead *TableRow, tbody *TableBody) *Table {
	return &Table{
		Class:     class,
		Id:        id,
		TableHead: thead,
		TableBody: tbody,
		Category:  "default",
	}
}

// NewDataTable ...
func NewDataTable(class string, id string, thead *TableRow, tbody *TableBody) *Table {
	return &Table{
		Class:     class,
		Id:        id,
		TableHead: thead,
		TableBody: tbody,
		Script:    &Script{Id: "#" + id, Load: true, Category: "dataTableUsers"},
		Category:  "default",
	}
}

// NewUsersTable ...
func NewUsersTable(users []*User) *Table {
	tId, _ := GenerateUuid()
	var tr []*TableRow
	for _, u := range users {
		tr = append(tr, NewUserRow(u))
	}
	return NewDataTable("display", tId, NewUserHeaderRow(), NewTableBody("", "", tr))
}
