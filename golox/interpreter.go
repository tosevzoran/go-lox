package golox

import "fmt"

// the interpreter struct needs to implement IExprVisitor interface
type Interpteter struct{}

func NewInterpreter() Interpteter {
	return Interpteter{}
}

func (i Interpteter) VisitBinaryExpr(expr BinaryExpr) (any, error) {
	left, err := i.evaluate(expr.left)

	if err != nil {
		return nil, err
	}

	right, err := i.evaluate(expr.right)

	if err != nil {
		return nil, err
	}

	switch expr.operator.tokenType {
	case MINUS:
		return left.(float64) - right.(float64), nil
	case SLASH:
		return left.(float64) / right.(float64), nil
	case STAR:
		return left.(float64) * right.(float64), nil
	case PLUS:
		// The + operator can also be used to concatenate two strings.
		if left, ok := left.(float64); ok {
			if right, ok := right.(float64); ok {
				return left + right, nil
			}
		}

		if left, ok := left.(string); ok {
			if right, ok := right.(string); ok {
				return left + right, nil
			}
		}

		return nil, fmt.Errorf("operands must be two numbers or two strings")
	case GREATER:
		return left.(float64) > right.(float64), nil
	case GREATER_EQUAL:
		return left.(float64) >= right.(float64), nil
	case LESS:
		return left.(float64) < right.(float64), nil
	case LESS_EQUAL:
		return left.(float64) <= right.(float64), nil
	case BANG_EQUAL:

		if err != nil {
			return nil, err
		}

		return !isEqual(left, right), nil
	case EQUAL_EQUAL:
		return isEqual(left, right), nil
	}

	return nil, fmt.Errorf("unsupported expression")
}

func (i Interpteter) VisitGroupingExpr(expr GroupingExpr) (any, error) {
	return i.evaluate(expr.expression)
}

func (Interpteter) VisitLiteralExpr(expr LiteralExpr) (any, error) {
	return expr.value, nil
}

func (i Interpteter) VisitUnaryExpr(expr UnaryExpr) (any, error) {
	right, err := i.evaluate(expr.right)

	if err != nil {
		return nil, err
	}

	switch expr.operator.tokenType {
	case MINUS:
		return -right.(float64), nil
	case BANG:
		return !isTruthy(right), nil
	}

	return nil, fmt.Errorf("unsupported expression")
}

func (i Interpteter) evaluate(expr IExpr) (any, error) {
	return expr.Accept(i)
}

func isTruthy(value any) bool {
	// false and nil are falsey, and everything else is truthy
	if value == nil {
		return false
	}

	if value == false {
		return false
	}

	return true

}

func isEqual(a any, b any) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a == b
}
