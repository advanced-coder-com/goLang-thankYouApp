package Db

import (
	"database/sql"
	handler "thankYou/Handler"

	_ "github.com/mattn/go-sqlite3"
)

// Model thankYou db contants. Cannot import thankYou/Model
// becaurse cycle imports
const TABLE_NAME string = "thank_you"
const ID_FIELD string = "id"
const LANGUAGE_FIELD string = "language"
const LANG_CODE_FIELD string = "lang_code"
const VALUE_FIELD string = "value"

func checkDb(db *sql.DB) {
	_, table_check := db.Query("select * from `" + TABLE_NAME + "`;")

	if table_check != nil {
		_, err := db.Exec("CREATE TABLE `" + TABLE_NAME + "` (`" + LANGUAGE_FIELD + "` INTEGER, `" + LANGUAGE_FIELD + "` TEXT, `" + LANG_CODE_FIELD + "` TEXT,`" + VALUE_FIELD + "` TEXT,	PRIMARY KEY(`id` AUTOINCREMENT))")
		handler.Handle(err)
		_, err = db.Exec("CREATE INDEX 'LANG_CODE' ON `thank_you` ('lang_code'	ASC);")
		handler.Handle(err)
	}
}
