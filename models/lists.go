package models

// NewList initializes a new list of groups for rendering
func NewList[T DataModel](m []T) *List {
	var listItems []*ListItem
	for _, gr := range m {
		listItems = append(listItems, NewListItem("group", gr.GetID(), gr.GetLabel()))
	}
	return NewUnorderedList("groups", "", listItems, nil)
}

// NewLinkedList initializes a new linked list of DataModel for rendering
func NewLinkedList[T DataModel](m []T, baseURL string, delete bool, search bool, chkBox bool) *List {
	var listItems []*ListItem
	listId, _ := GenerateUuid()
	for _, gr := range m {
		gLink := NewLink("pill", "", baseURL+gr.GetClass(true)+"/"+gr.GetID(), gr.GetLabel(), false)
		var ops []*ItemOption
		if delete {
			defForm := InitializePopupDeleteForm(gr)
			btn := defForm.Popup
			defForm.Popup = nil
			delOp := NewDeleteOption(defForm, btn)
			ops = append(ops, delOp)
		}
		if chkBox {
			reqURL := baseURL + gr.GetClass(true) + "/" + gr.GetID() + "/check"
			chkInput := NewCheckboxInput("", "pill checkbox", "check"+gr.GetID(), gr.GetBoolField("Completed"), reqURL)
			chkOp := NewCheckOption("", chkInput)
			ops = append(ops, chkOp)
		}
		listItems = append(listItems, NewLinkListItem("pill", gr.GetID(), gLink, ops))
	}
	if search {
		inId, _ := GenerateUuid()
		searchInput := NewSearchInput("linked filter", inId, listId, "Enter a name: ")
		return NewUnorderedList("linked", listId, listItems, searchInput)
	}
	return NewUnorderedList("linked", listId, listItems, nil)
}
