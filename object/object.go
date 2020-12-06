package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Objects name's for ObjectType
const (
	NULL_OBJ  = "NULL"
	ERROR_OBJ = "ERROR"

	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"

	RETURN_VALUE_OBJ = "RETURN_VALUE"

	FUNCTION_OBJ = "FUNCTION"
)

// Types
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}
