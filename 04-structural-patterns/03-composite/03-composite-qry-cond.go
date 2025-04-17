package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	ToSQL() string
}

type Condition struct {
	Field, Operator, Value string
}

func (c *Condition) ToSQL() string {
	return fmt.Sprintf("%s %s '%s'", c.Field, c.Operator, c.Value)
}

type AndExpression struct {
	Children []Expression
}

func (a *AndExpression) ToSQL() string {
	var parts []string
	for _, child := range a.Children {
		parts = append(parts, child.ToSQL())
	}
	return "(" + strings.Join(parts, " AND ") + ")"
}

type OrExpression struct {
	Children []Expression
}

func (o *OrExpression) ToSQL() string {
	var parts []string
	for _, child := range o.Children {
		parts = append(parts, child.ToSQL())
	}
	return "(" + strings.Join(parts, " OR ") + ")"
}

func main() {
	cond1 := &Condition{"age", ">", "30"}
	cond2 := &Condition{"city", "=", "Delhi"}
	cond3 := &Condition{"salary", ">", "50000"}

	andGroup := &AndExpression{[]Expression{cond1, cond2}}
	orGroup := &OrExpression{[]Expression{andGroup, cond3}}

	fmt.Println("WHERE", orGroup.ToSQL())
}
