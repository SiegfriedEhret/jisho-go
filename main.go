package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
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

var opts struct {
	Db   string `long:"db" description:"Path to the kanjidb.sqlite file" default:"./kanjidb.sqlite"`
	Args struct {
		Id     string
		Kanjis []string
	} `positional-args:"yes" required:"yes"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		logrus.WithError(err).Fatal("Failed to read program arguments")
	}

	logrus.WithField("Db", opts.Db).Debug("Database path")
	logrus.WithField("Kanjis", opts.Args.Kanjis).Debug("Kanjis to search")

	fmt.Println("jisho 辞書")

	if len(opts.Args.Kanjis) == 0 {
		log.Fatal("Please give me a kanji!")
	}

	c := make(chan KanjiData, len(opts.Args.Kanjis))
	database, err := sql.Open("sqlite3", opts.Db)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	for i := 0; i < len(opts.Args.Kanjis); i++ {
		go search(database, opts.Args.Kanjis[i], c)
	}

	count := 0
	for i := range c {
		fmt.Println(i)

		count++
		if count == len(opts.Args.Kanjis) {
			return
		}
	}
}

func search(database *sql.DB, kanji string, c chan KanjiData) {
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

	rows, err := database.Query(query)
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
