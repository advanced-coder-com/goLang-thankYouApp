package ThankYouModel

import (
	"math/rand"
	"thankYou/Db"
	handler "thankYou/Handler"
	"time"
)

type ThankYou struct {
	id        int
	language  string
	lang_code string
	value     string
}

const TABLE_NAME string = "thank_you"

func GetList(conditions []Db.Condition) []ThankYou {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	rows := DbAdapter.GetList(TABLE_NAME, conditions)

	var result []ThankYou
	for rows.Next() {
		var thankYou ThankYou = ThankYou{}
		err := rows.Scan(&thankYou.id, &thankYou.language, &thankYou.lang_code, &thankYou.value)
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
	err := row.Scan(&thankYou.id, &thankYou.language, &thankYou.lang_code, &thankYou.value)
	handler.Handle(err)
	return thankYou
}

// func getAdapter() Db.DbAdapter {
// 	DbAdapter := Db.NewDbAdapter()
// 	defer DbAdapter.CloseDb()
// 	return DbAdapter
// }

func GetRandom() ThankYou {
	// DbAdapter := getAdapter()
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	count := DbAdapter.GetCount(TABLE_NAME)
	rand.Seed(time.Now().UnixNano())
	randomId := rand.Intn(count) + 1
	return GetById(randomId)
}

func GetCount() int {
	DbAdapter := Db.NewDbAdapter()
	defer DbAdapter.CloseDb()
	return DbAdapter.GetCount(TABLE_NAME)
}

func (subject ThankYou) ToArray() []string {
	return []string{subject.value, subject.language, subject.lang_code}
}
