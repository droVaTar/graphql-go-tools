package ast

import "github.com/drovatar/graphql-go-tools/pkg/lexer/position"

type SchemaExtension struct {
	ExtendLiteral position.Position
	SchemaDefinition
}
