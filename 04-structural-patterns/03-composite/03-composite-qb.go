package main

import (
	"fmt"
	"strings"
)

// Expression represents part of a WHERE clause.
type Expression interface {
	ToSQL() string
}

type Condition struct {
	Field, Operator, Value string
}

func (c *Condition) ToSQL() string {
	return fmt.Sprintf("%s %s '%s'", c.Field, c.Operator, c.Value)
}

type And struct {
	Expressions []Expression
}

func (a *And) ToSQL() string {
	parts := make([]string, len(a.Expressions))
	for i, expr := range a.Expressions {
		parts[i] = expr.ToSQL()
	}
	return "(" + strings.Join(parts, " AND ") + ")"
}

type Or struct {
	Expressions []Expression
}

func (o *Or) ToSQL() string {
	parts := make([]string, len(o.Expressions))
	for i, expr := range o.Expressions {
		parts[i] = expr.ToSQL()
	}
	return "(" + strings.Join(parts, " OR ") + ")"
}

// QueryBuilder builds SELECT queries using the expression tree.
type QueryBuilder struct {
	Table     string
	Columns   []string
	Condition Expression
}

func (qb *QueryBuilder) Build() string {
	cols := "*"
	if len(qb.Columns) > 0 {
		cols = strings.Join(qb.Columns, ", ")
	}
	sql := fmt.Sprintf("SELECT %s FROM %s", cols, qb.Table)
	if qb.Condition != nil {
		sql += " WHERE " + qb.Condition.ToSQL()
	}
	return sql + ";"
}

func main() {
	// Example
	cond1 := &Condition{"age", ">", "30"}
	cond2 := &Condition{"city", "=", "Delhi"}
	cond3 := &Condition{"salary", ">", "50000"}

	expr := &Or{
		Expressions: []Expression{
			&And{[]Expression{cond1, cond2}},
			cond3,
		},
	}

	qb := &QueryBuilder{
		Table:     "employees",
		Columns:   []string{"id", "name", "salary"},
		Condition: expr,
	}

	fmt.Println(qb.Build())
}
