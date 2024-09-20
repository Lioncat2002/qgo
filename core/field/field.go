package field

import "fmt"

type Field struct {
	Name      string
	Unique    bool
	Nullable  bool
	Default   interface{}
	FieldType string
}

func Int(name string) Field {
	return Field{Name: name, Unique: false, Nullable: false, Default: nil, FieldType: "int"}
}

func String(name string, size uint64) Field {
	fieldType := fmt.Sprintf("varchar(%d)", size)
	_ = fieldType
	return Field{Name: name, Unique: false, Nullable: false, Default: nil, FieldType: "text"}
}

func (field Field) SetUnique() Field {
	field.Unique = true
	return field
}

func (field Field) SetDefault(value interface{}) Field {
	field.Default = value
	return field
}

func (field Field) SetNullable() Field {
	field.Nullable = true
	return field
}
