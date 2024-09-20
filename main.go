package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Lioncat2002/qgo/core/field"
	"github.com/Lioncat2002/qgo/core/qgo"
)

type User struct {
	qgo.Schema
}

func (User) Fields() []field.Field {
	return []field.Field{
		field.Int("id").SetUnique(),
		field.String("name", 255).SetDefault("unknown"),
	}
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln("failed to create db", err)
	}
	err = qgo.Migrate(db, User{})
	if err != nil {
		log.Fatalln("failed to create table user", err)
	}
}
