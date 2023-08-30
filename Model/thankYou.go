package ThankYouModel

import (
	"fmt"
	"strconv"
	"thankYou/Db"
	handler "thankYou/Handler"
)

type ThankYou struct {
	Id        *int
	Language  string
	Lang_code string
	Value     string
}

const TABLE_NAME string = "thank_you"
const ID_FIELD string = "id"
const LANGUAGE_FIELD string = "language"
const LANG_CODE_FIELD string = "lang_code"
const VALUE_FIELD string = "value"

func (subject ThankYou) Save() bool {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()

	var err error
	if insertOrUpdate(subject) {
		err = DbAdapter.Insert(TABLE_NAME, prepareValuesForSql(subject))
	} else {
		conditions := []Db.Condition{prepareIdWhereCondition(subject)}
		err = DbAdapter.Update(TABLE_NAME, prepareValuesForSql(subject), conditions)
	}
	handler.Handle(err)
	return err == nil
}

func Delete(langCode string) bool {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	subject := GetByLangCode(langCode)
	conditions := []Db.Condition{prepareIdWhereCondition(subject)}

	err := DbAdapter.Delete(TABLE_NAME, conditions)
	handler.Handle(err)
	return err == nil
}

func GetList(conditions []Db.Condition) []ThankYou {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	rows := DbAdapter.GetList(TABLE_NAME, conditions)

	var result []ThankYou
	for rows.Next() {
		var thankYou ThankYou = ThankYou{}
		err := rows.Scan(&thankYou.Id, &thankYou.Language, &thankYou.Lang_code, &thankYou.Value)
		handler.Handle(err)
		result = append(result, thankYou)
	}
	return result
}

func GetById(id int) ThankYou {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	row := DbAdapter.GetById(TABLE_NAME, id)
	var thankYou ThankYou = ThankYou{}
	err := row.Scan(&thankYou.Id, &thankYou.Language, &thankYou.Lang_code, &thankYou.Value)
	handler.Handle(err)
	return thankYou
}

func GetByLangCode(langCode string) ThankYou {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	row := DbAdapter.GetByField(TABLE_NAME, LANG_CODE_FIELD, langCode)
	var thankYou ThankYou = ThankYou{}
	err := row.Scan(&thankYou.Id, &thankYou.Language, &thankYou.Lang_code, &thankYou.Value)
	handler.Handle(err)
	return thankYou
}

func GetRandom() ThankYou {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	var thankYou ThankYou = ThankYou{}
	row := DbAdapter.GetRandom(TABLE_NAME)
	err := row.Scan(&thankYou.Id, &thankYou.Language, &thankYou.Lang_code, &thankYou.Value)
	handler.Handle(err)
	return thankYou
}

func (subject ThankYou) ToArray() []string {
	return []string{subject.Value, subject.Language, subject.Lang_code, subject.Value + " (" + subject.Language + ")"}
}

func insertOrUpdate(subject ThankYou) bool {
	return subject.Id == nil
}

func prepareValuesForSql(subject ThankYou) [][]string {
	var result [][]string
	if subject.Id != nil {
		fmt.Println(subject.Id)
		result = append(result, []string{ID_FIELD, strconv.Itoa(*subject.Id)})
	}
	result = append(result,
		[]string{LANGUAGE_FIELD, addQuotes(subject.Language)},
		[]string{LANG_CODE_FIELD, addQuotes(subject.Lang_code)},
		[]string{VALUE_FIELD, addQuotes(subject.Value)},
	)
	return result
}

func addQuotes(input string) string {
	return "'" + input + "'"
}

func prepareIdWhereCondition(subject ThankYou) Db.Condition {
	return Db.Condition{Column: "id", Value: strconv.Itoa(*subject.Id), Predicate: "=", LogicOperator: "AND"}
}
