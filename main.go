package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// KanjiData represents a kanji
type KanjiData struct {
	Kanji    string
	Meaning  string
	Onyomi   string
	Kunyomi  string
	Elements string
	PartOf   string
	Jlpt     string
	Strokes  int
}

func main() {
	fmt.Println("jisho", os.Args)

	db, err := sql.Open("sqlite3", "./data/kanjidb.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	KanjiData := search(db, "冥")

	fmt.Println(KanjiData)

	defer db.Close()
}

func search(db *sql.DB, kanji string) KanjiData {
	query := `SELECT
			E.kanji,
			E.strokes,
			K.meaning,
			onyomi,
			kunyomi,
			elements,
			part_of,
			jlpt
		FROM elements AS E
		JOIN kanjidict AS K ON E.kanji = K.kanji
		WHERE E.kanji = '` + kanji + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var kanjiData KanjiData

	for rows.Next() {
		var kanji, meaning, onyomi, kunyomi, elements, partOf, jlpt string
		var strokes int

		rows.Scan(&kanji, &strokes, &meaning, &onyomi, &kunyomi, &elements, &partOf, &jlpt)

		kanjiData = KanjiData{kanji, meaning, onyomi, kunyomi, elements, partOf, jlpt, strokes}

	}

	return kanjiData
}

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
