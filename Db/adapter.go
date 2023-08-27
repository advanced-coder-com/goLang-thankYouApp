package Db

import (
	"database/sql"
	"fmt"
	"strconv"
	handler "thankYou/Handler"
)

type DbAdapter struct {
	db *sql.DB
}

// Column - colunm name in where
// Predicate - Simple predicates use one of the operators =, <>, >, >=, <, <=, IN, BETWEEN, LIKE, IS NULL or IS NOT NULL.
// Value - value for comparison
// logicOperator - OR / AND
type Condition struct {
	Column        string
	Value         string
	Predicate     string
	LogicOperator string
}

func NewDbAdapter() DbAdapter {
	var err error
	subject := DbAdapter{}
	subject.db, err = sql.Open("sqlite3", DB_NAME)
	handler.Handle(err)
	checkDb(subject.db)
	return subject
}

func (DbAdapter DbAdapter) GetById(tableName string, id int) *sql.Row {
	var sql string
	sql += getSelect(tableName)
	conditions := []Condition{{Column: "id", Value: strconv.Itoa(id), Predicate: "="}}
	sql += getWhere(conditions)
	return DbAdapter.db.QueryRow(sql)
}

func (DbAdapter DbAdapter) GetList(tableName string, conditions []Condition) *sql.Rows {
	var sql string
	sql += getSelect(tableName)
	sql += getWhere(conditions)
	rows, err := DbAdapter.db.Query(sql)
	if err != nil {
		rows.Close()
	}
	handler.Handle(err)

	return rows
}

func (DbAdapter DbAdapter) GetCount(tableName string) int {
	var cnt int
	_ = DbAdapter.db.QueryRow(`select count(*) from ` + tableName).Scan(&cnt)
	return cnt
}

func (DbAdapter DbAdapter) GetDb() *sql.DB {

	return DbAdapter.db
}

func (DbAdapter DbAdapter) CloseDb() {
	DbAdapter.db.Close()
}

func getSelect(tableName string) string {
	return fmt.Sprintf("SELECT * FROM %s", tableName)
}

func getWhere(conditions []Condition) string {
	if len(conditions) == 0 {
		return ""
	}
	var result = " WHERE"
	for index, condition := range conditions {
		var row string
		if index > 0 {
			row += fmt.Sprintf(" %s", condition.LogicOperator)
		}

		row += fmt.Sprintf(" %s%s%s", condition.Column, condition.Predicate, condition.Value)
		result += row
	}
	return result
}
