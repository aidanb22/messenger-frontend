package models

// Field ...
type Field struct {
	Name     string
	Class    string
	Id       string
	Type     string
	Label    string
	Value    string
	Category string
}

// LoadFields ...
func LoadFields(fieldStrs [][]string) []Field {
	var fields []Field
	// Name, Class, Id, Type, Label, DefaultVal
	for _, f := range fieldStrs {
		newField := Field{f[0], f[1], f[2], f[3], f[4], f[5], f[5]}
		fields = append(fields, newField)
	}
	return fields
}

/*
func (fm *Field) Update() Auth {
	auth := Auth{}
	auth.Register(fm.Name, fm.Class, fm.Id, fm.Type, fm.Label)
	return auth
}
*/
/*
func (f *Form) NewForm(auth Auth, fields []Field, button Button) Form {
	return Form{Fields: fields, Button: button}
}
*/
/*
func (bm *Button) Update() Auth {
	auth := Auth{}
	auth.Register(bm.Name, bm.Class, bm.Id, bm.Type, bm.Label)
	return auth
}
*/
