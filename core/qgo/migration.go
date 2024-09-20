package qgo

import (
	"bytes"
	"database/sql"
	"log"
	"reflect"
	"text/template"

	"github.com/Lioncat2002/qgo/core"
	"github.com/Lioncat2002/qgo/core/field"
)

type Table struct {
	Name   string
	Fields []field.Field
}

func createTable(db *sql.DB, template template.Template, table Table) error {
	var query bytes.Buffer

	err := template.Execute(&query, table)
	if err != nil {
		return err
	}
	log.Println(query.String())
	result, err := db.Exec(query.String())
	if err != nil {
		return err
	}
	log.Println("Migration Result:", result)
	return nil
}

func Migrate(db *sql.DB, schema SchemaInterface) error {
	t := reflect.TypeOf(schema)
	tmpl, err := template.New("createTable").Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b }, // Helper function for subtraction
		"len": func(a []field.Field) int { return len(a) },
	}).Parse(core.CREATE_TABLE)
	if err != nil {
		return err
	}

	table := Table{
		Name:   t.Name(),
		Fields: schema.Fields(),
	}
	if err := createTable(db, *tmpl, table); err != nil {
		return err
	}
	return nil
}
