package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("jisho", os.Args)

	db, err := sql.Open("sqlite3", "./data/kanjidb.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

/*
func search(db, kanji) {

}
*/

/*
function search(kanji, cb) {
  doSearch(`SELECT *
    FROM elements AS E
    JOIN kanjidict AS K ON E.kanji = K.kanji
    WHERE E.kanji = '${kanji}'`, cb);
}

function searchSentence(kanji, cb) {
  doSearch(`SELECT *
    FROM sentences
    WHERE kanji LIKE '%${kanji}%';`, cb);
}

function doSearch(query, cb) {
  debug('Query:', query);

  db.all(query, cb);
}
*/
