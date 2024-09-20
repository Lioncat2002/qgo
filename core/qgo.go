package core

import "database/sql"

type Client struct {
	DB *sql.DB
}
