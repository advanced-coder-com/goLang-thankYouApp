package Db

import "fmt"

// Method helps to build SQL statement

// Prefixes getters
func getSelect(tableName string) string {
	return fmt.Sprintf("SELECT * FROM %s", tableName)
}

func getInsert(tableName string) string {
	return fmt.Sprintf("INSERT INTO %s", tableName)
}

func getUpdate(tableName string) string {
	return fmt.Sprintf(" UPDATE %s", tableName)
}

func getDelete(tableName string) string {
	return fmt.Sprintf("DELETE FROM %s", tableName)
}

// Where generating
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

// Insert helpers
func getInsertValues(data [][]string) string {
	var columns string = "("
	var values string = "("
	for index, field := range data {
		if index > 0 {
			columns += ", "
			values += ", "
		}
		columns += field[0]
		values += field[1]
	}
	columns += ")"
	values += ")"
	return columns + " VALUES " + values
}

// Update helpers
func getUpdateValues(data [][]string) string {
	var result string = " SET "
	for index, field := range data {
		if index > 0 {
			result += ", "
		}
		result = fmt.Sprintf("%s=%s", field[0], field[1])
	}

	return result
}
