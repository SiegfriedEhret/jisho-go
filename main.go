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
	JLPT     string
	Strokes  int
}

func (k KanjiData) String() string {
	return fmt.Sprintf(`
		Kanji %s

	Onyomi %s
	Kunyomi %s
	Strokes %d
	Meaning %s
	JLPT %s
	Parts %s
	Part of %s
	`, k.Kanji, k.Onyomi, k.Kunyomi, k.Strokes, k.Meaning, k.JLPT, k.Elements, k.PartOf)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Please give me a kanji!")
	}

	c := make(chan KanjiData, len(args))
	db, err := sql.Open("sqlite3", "./data/kanjidb.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(args); i++ {
		go search(db, args[i], c)
	}

	count := 0
	for i := range c {
		fmt.Println(i)

		count++
		if count == len(args) {
			return
		}
	}

	defer db.Close()
}

func search(db *sql.DB, kanji string, c chan KanjiData) {
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

	c <- kanjiData
}

/*
From https://gitlab.com/SiegfriedEhret/jisho

TODO:

function searchSentence(kanji, cb) {
  doSearch(`SELECT *
    FROM sentences
    WHERE kanji LIKE '%${kanji}%';`, cb);
}

*/
