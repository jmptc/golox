package exprs

import "github.com/jmptc/golox/token"

type Expr interface {
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type Grouping struct {
	Expr Expr
}

type Literal struct {
	Value string
}

type Unary struct {
	Operator token.Token
	Right    Expr
}
