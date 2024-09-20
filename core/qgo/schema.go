package qgo

import (
	"github.com/Lioncat2002/qgo/core/field"
)

type SchemaInterface interface {
	Fields() []field.Field
}

type Schema struct {
	SchemaInterface
}
