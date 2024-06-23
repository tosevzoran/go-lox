package golox

type IStmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) error
	VisitPrintStmt(stmt PrintStmt) error
}

type IStmt interface {
	Accept(v IStmtVisitor) error
}

type ExpressionStmt struct {
	expr IExpr
}

func NewExpressionStmt(expr IExpr) ExpressionStmt {
	return ExpressionStmt{expr}
}

func (s ExpressionStmt) Accept(v IStmtVisitor) error {
	return v.VisitExpressionStmt(s)
}

type PrintStmt struct {
	expr IExpr
}

func NewPrintStmt(expr IExpr) PrintStmt {
	return PrintStmt{expr}
}

func (p PrintStmt) Accept(v IStmtVisitor) error {
	return v.VisitPrintStmt(p)
}
