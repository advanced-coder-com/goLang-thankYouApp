package Db

import (
	"database/sql"
	"os"
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
		_, err := db.Exec("CREATE TABLE `" + TABLE_NAME + "` (`" + ID_FIELD + "` INTEGER, `" + LANGUAGE_FIELD + "` TEXT, `" + LANG_CODE_FIELD + "` TEXT,`" + VALUE_FIELD + "` TEXT,	PRIMARY KEY(`id` AUTOINCREMENT))")
		handler.Handle(err)
		_, err = db.Exec("CREATE UNIQUE INDEX 'LANG_CODE' ON `thank_you` ('lang_code'	ASC);")
		handler.Handle(err)
		_, err = db.Exec(`INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('grazie', 'Italian', 'It');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('thank you', 'English', 'En');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('cảm ơn bạn', 'Vietnam', 'Vi');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('დიდი მადლობა', 'Georgian', 'Ge');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('gracias', 'Spanish', 'Es');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('謝謝', 'Chinese', 'Zh');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('paldies', 'Latvian', 'Lv');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('dank u wel', 'Netherlands', 'NL');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('danke', 'German', 'De');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('dziekuje', 'Poland', 'Pl');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('mulţumesc', 'Romanian', 'Ro');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('ขอขอบคุณ', 'Thai', 'Th');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('iyiyim, teşekkürler', 'Turkish', 'Tr');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('köszönöm', 'Hungarian', 'Hu');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('ありがとう', 'Japanese', 'Ja');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('tack', 'Sweden', 'Se');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('hvala', 'Croatian', 'Hr');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('kiitos', 'Finnish', 'Fi');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('salamat', 'Filipino', 'Fil');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('рақмет сізге', 'Kazakh', 'Kk');
		INSERT INTO "thank_you" ("value", "language", "lang_code") VALUES ('дякую', 'Ukrainian', 'Uk');`)
		handler.Handle(err)
	}
}

func checkDbDirectory(dirName string) {
	_, err := os.Stat("./" + dirName)
	if err != nil {
		err := os.Mkdir(dirName, 0755)
		handler.Handle(err)
	}
}
