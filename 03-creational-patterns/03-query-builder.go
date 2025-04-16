package main

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	selectCols []string
	table      string
	where      string
	orderBy    string
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(cols ...string) *QueryBuilder {
	qb.selectCols = cols
	return qb
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.table = table
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.where = condition
	return qb
}

func (qb *QueryBuilder) OrderBy(order string) *QueryBuilder {
	qb.orderBy = order
	return qb
}

func (qb *QueryBuilder) Build() string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	if len(qb.selectCols) > 0 {
		sb.WriteString(strings.Join(qb.selectCols, ", "))
	} else {
		sb.WriteString("*")
	}
	sb.WriteString(" FROM ")
	sb.WriteString(qb.table)
	if qb.where != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(qb.where)
	}
	if qb.orderBy != "" {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(qb.orderBy)
	}
	sb.WriteString(";")
	return sb.String()
}

func main() {
	query := NewQueryBuilder().
		Select("id", "name").
		From("users").
		Where("active = 1").
		OrderBy("name").
		Build()

	fmt.Println(query)
}
