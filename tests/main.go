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
		field.Int("Id").SetUnique(),
		field.String("Name", 255).SetDefault("unknown"),
	}
}

type UserSchema struct {
	Id   int
	Name string
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
	result, err := qgo.Insert(User{}).Columns(UserSchema{}).Values(UserSchema{Id: 5, Name: "meow"}).Exec(db)
	log.Println(result, err)
	result, err = qgo.Select(UserSchema{}).From(User{}).Exec(db)
	log.Println(result, err)
	data := UserSchema{}
	for result.(*sql.Rows).Next() {
		err := result.(*sql.Rows).Scan(&data.Id, &data.Name)
		if err != nil {
			log.Fatal(err)
		}

		// Output the data
		log.Println(data)
	}
}
