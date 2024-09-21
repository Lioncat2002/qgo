package qgo

import "database/sql"

type Client struct {
	DB *sql.DB
}
