package Db

import (
	"database/sql"
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

// Constructor
func NewDbAdapter() DbAdapter {
	var err error
	subject := DbAdapter{}
	checkDbDirectory(DB_DIR_NAME)
	subject.db, err = sql.Open("sqlite3", DB_NAME)
	handler.Handle(err)
	checkDb(subject.db)
	return subject
}

// Get record by Id (int)
func (DbAdapter DbAdapter) GetById(tableName string, id int) *sql.Row {
	var sql string
	sql += getSelect(tableName)
	conditions := []Condition{{Column: "id", Value: strconv.Itoa(id), Predicate: "="}}
	sql += getWhere(conditions)
	return DbAdapter.db.QueryRow(sql)
}

// Get record by string value
func (DbAdapter DbAdapter) GetByField(tableName string, fieldName string, fieldValue string) *sql.Row {
	var sql string
	sql += getSelect(tableName)
	conditions := []Condition{{Column: fieldName, Value: "'" + fieldValue + "'", Predicate: "="}}
	sql += getWhere(conditions)
	return DbAdapter.db.QueryRow(sql)
}

// Get random record
func (DbAdapter DbAdapter) GetRandom(tableName string) *sql.Row {
	var sql string
	sql += getSelect(tableName)
	sql += " ORDER BY RANDOM() LIMIT 1"
	return DbAdapter.db.QueryRow(sql)
}

// Get list of records by conditions
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

func (DbAdapter DbAdapter) Insert(tableName string, values [][]string) error {
	var sql string
	sql += getInsert(tableName)
	sql += getInsertValues(values)
	_, err := DbAdapter.db.Exec(sql)
	return err
}

func (DbAdapter DbAdapter) Update(tableName string, values [][]string, conditions []Condition) error {
	var sql string
	sql += getUpdate(tableName)
	sql += getUpdateValues(values)
	sql += getWhere(conditions)
	_, err := DbAdapter.db.Exec(sql)
	return err
}

func (DbAdapter DbAdapter) Delete(tableName string, conditions []Condition) error {
	var sql string
	sql += getDelete(tableName)
	sql += getWhere(conditions)
	_, err := DbAdapter.db.Exec(sql)
	return err
}

// Destructor
func (DbAdapter DbAdapter) CloseDb() {
	DbAdapter.db.Close()
}
