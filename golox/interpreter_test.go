package golox

import "testing"

type expressionTest struct {
	name       string
	expression IExpr
	expected   any
}

func TestUnaryExpressionEval(t *testing.T) {
	unaryExpressionsTests := []expressionTest{
		{
			name:       "-3",
			expression: NewUnaryExpr(NewToken(MINUS, "-", "-", 0), NewLiteralExpr(float64(3))),
			expected:   -3.0,
		},
		{
			name:       "!true",
			expression: NewUnaryExpr(NewToken(BANG, "!", "!", 0), NewLiteralExpr(true)),
			expected:   false,
		},
		{
			name:       "!nil",
			expression: NewUnaryExpr(NewToken(BANG, "!", "!", 0), NewLiteralExpr(nil)),
			expected:   true,
		},
	}

	runExprTests(t, unaryExpressionsTests)
}

func TestBinaryExpressionEval(t *testing.T) {

	binaryExpressionsTests := []expressionTest{
		{
			name:       "2-3=-1",
			expression: NewBinaryExpr(NewLiteralExpr(float64(2)), NewToken(MINUS, "-", "-", 0), NewLiteralExpr(float64(3))),
			expected:   -1.0,
		},
		{
			name: "2-3>9=false",
			expression: NewBinaryExpr(
				NewBinaryExpr(
					NewLiteralExpr(float64(2)),
					NewToken(MINUS, "-", "-", 0),
					NewLiteralExpr(float64(3)),
				),
				NewToken(GREATER, ">", ">", 0),
				NewLiteralExpr(float64(9)),
			),
			expected: false,
		},
		{
			name: "6-8*9=-66",
			expression: NewBinaryExpr(
				NewLiteralExpr(float64(6)),
				NewToken(MINUS, "-", "-", 0),
				NewBinaryExpr(
					NewLiteralExpr(float64(8)),
					NewToken(STAR, "*", "*", 0),
					NewLiteralExpr(float64(9)),
				),
			),
			expected: -66.0,
		},
		{
			name: "(6-8)*9=-18",
			expression: NewBinaryExpr(
				NewGroupingExpr(
					NewBinaryExpr(
						NewLiteralExpr(float64(6)),
						NewToken(MINUS, "-", "-", 0),
						NewLiteralExpr(float64(8)),
					),
				),
				NewToken(STAR, "*", "*", 0),
				NewLiteralExpr(float64(9)),
			),
			expected: -18.0,
		},
	}

	runExprTests(t, binaryExpressionsTests)
}

func runExprTests(t *testing.T, expressions []expressionTest) {
	t.Helper()

	loxInterpreter := NewInterpreter()

	for _, tt := range expressions {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := loxInterpreter.evaluate(tt.expression)

			if got != tt.expected {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})
	}
}
