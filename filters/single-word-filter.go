package filters

import (
	"database/sql"
	"fmt"
	"profanityfilteritaapi/structures"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type FilterResult struct {
	Result  bool
	Details []Profanity
}

type Profanity struct {
	Id          int
	Profanity   string
	IsRacist    bool
	IsReligious bool
	IsSexist    bool
}

func SingleWordFilter(word string) FilterResult {
	s := structures.Settings{}.InitSettings()
	fmt.Println("Db file path:", s.Dblocation)
	db, err := sql.Open("sqlite3", s.Dblocation)
	if err != nil {
		panic(err)
	}

	word = strings.ToLower(word)
	newWord := ""
	switches := map[string]string{
		"a": "4àáâãäåæ",
		"e": "3èéêë",
		"i": "1ìíîï",
		"o": "0òóôõö",
		"u": "ùúûü",
	}

	for w := range word {
		stringedw := string(word[w])
		for k := range switches {
			if In(stringedw, switches[k]) {
				newWord += k
				continue
			}
		}
		newWord += stringedw
	}

	qr, err := db.Prepare("SELECT * FROM profanities WHERE profanity=?;")
	if err != nil {
		panic(err)
	}
	res, err := qr.Query(word)
	if err != nil {
		panic(err)
	}
	var collector []Profanity
	for res.Next() {
		var temp Profanity
		res.Scan(&temp.Id, &temp.Profanity, &temp.IsRacist, &temp.IsReligious, &temp.IsSexist)
		collector = append(collector, temp)
	}
	out := FilterResult{}
	out.Details = collector
	if len(collector) != 0 {
		out.Result = true
	} else {
		out.Result = false
	}
	return out
}
