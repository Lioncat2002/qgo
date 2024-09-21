package qgo

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type Query struct {
	query  string
	values []interface{}
}

func Select(columns interface{}) Query {
	query := Query{query: "", values: nil}
	t := reflect.TypeOf(columns)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.String {
		query.query = fmt.Sprintf("SELECT %s ", columns.(string))
		return query
	}

	fieldNames := t.Field(0).Name
	for i := 1; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldNames = fmt.Sprintf("%s , %s", fieldNames, field.Name)
	}
	query.query = fmt.Sprintf("SELECT %s ", fieldNames)

	return query
}

func (query Query) From(table interface{}) Query {
	t := reflect.TypeOf(table)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	query.query = fmt.Sprintf("%s FROM %s ; ", query.query, t.Name())
	return query
}

func Insert(table interface{}) Query {
	query := Query{query: "", values: nil}
	t := reflect.TypeOf(table)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	query.query = fmt.Sprintf("INSERT INTO %s", t.Name())
	return query
}

func (query Query) Columns(columns interface{}) Query {
	t := reflect.TypeOf(columns)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	tmpl := t.Field(0).Name
	for i := 1; i < t.NumField(); i++ {
		tmpl = fmt.Sprintf("%s , %s ", tmpl, t.Field(i).Name)
	}
	query.query = fmt.Sprintf("%s ( %s ) ", query.query, tmpl)
	return query
}

func (query Query) Values(values interface{}) Query {
	t := reflect.TypeOf(values)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	tmpl := "?"
	data := make([]interface{}, t.NumField())
	data[0] = reflect.ValueOf(values).Field(0).Interface()
	for i := 1; i < t.NumField(); i++ {
		tmpl = fmt.Sprintf("%s , ? ", tmpl)
		data[i] = reflect.ValueOf(values).Field(i).Interface()
	}
	query.query = fmt.Sprintf("%s VALUES ( %s ) ", query.query, tmpl)

	query.values = data
	return query
}
func (query Query) Exec(db *sql.DB) (interface{}, error) {
	if query.values != nil {
		log.Println(query.query, query.values)
		result, err := db.Exec(query.query, query.values...)
		return result, err
	}
	result, err := db.Query(query.query)
	return result, err
}
