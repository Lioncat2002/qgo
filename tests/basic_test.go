package tests

import (
	"database/sql"
	"testing"

	"github.com/Lioncat2002/qgo/core/field"
	"github.com/Lioncat2002/qgo/core/qgo"
	_ "github.com/mattn/go-sqlite3"
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

func TestUserCreate(t *testing.T) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal("failed to create db", err)
	}
	err = qgo.Migrate(db, User{})
	if err != nil {
		t.Fatal("failed to create db", err)
	}
}
