package golox

type IExprVisitor interface {
	VisitLiteralExpr(expr LiteralExpr) (any, error)
	VisitBinaryExpr(expr BinaryExpr) (any, error)
	VisitGroupingExpr(expr GroupingExpr) (any, error)
	VisitUnaryExpr(expr UnaryExpr) (any, error)
}

type IExpr interface {
	Accept(v IExprVisitor) (any, error)
}

/*
Each expression gets it's own struct

expression     → literal
               | unary
               | binary
               | grouping ;

literal        → NUMBER | STRING | "true" | "false" | "nil" ;
grouping       → "(" expression ")" ;
unary          → ( "-" | "!" ) expression ;
binary         → expression operator expression ;
operator       → "==" | "!=" | "<" | "<=" | ">" | ">=" | "+"  | "-"  | "*" | "/" ;

*/

type LiteralExpr struct {
	value any
}

func NewLiteralExpr(v any) LiteralExpr {
	return LiteralExpr{v}
}

func (expr LiteralExpr) Accept(v IExprVisitor) (any, error) {
	return v.VisitLiteralExpr(expr)
}

type BinaryExpr struct {
	left     IExpr
	operator Token
	right    IExpr
}

func NewBinaryExpr(l IExpr, o Token, r IExpr) BinaryExpr {
	return BinaryExpr{l, o, r}
}

func (expr BinaryExpr) Accept(v IExprVisitor) (any, error) {
	return v.VisitBinaryExpr(expr)
}

type GroupingExpr struct {
	expression IExpr
}

func NewGroupingExpr(e IExpr) GroupingExpr {
	return GroupingExpr{e}
}

func (expr GroupingExpr) Accept(v IExprVisitor) (any, error) {
	return v.VisitGroupingExpr(expr)
}

type UnaryExpr struct {
	operator Token
	right    IExpr
}

func NewUnaryExpr(o Token, r IExpr) UnaryExpr {
	return UnaryExpr{o, r}
}

func (expr UnaryExpr) Accept(v IExprVisitor) (any, error) {
	return v.VisitUnaryExpr(expr)
}
