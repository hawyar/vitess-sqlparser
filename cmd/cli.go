package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

type Payload struct {
	Data string `json:"data"`
	Err  string `json:"err"`
}

type Create struct {
	Table   string            `json:"table"`
	Columns map[string]string `json:"columns"`
}

func main() {

	p := Payload{
		Data: "",
		Err:  "",
	}

	var destination string

	flag.StringVar(&destination, "d", "destination", "Speciify location of output")
	flag.Parse()

	if len(os.Args) < 2 {
		p.Err = "No query specified"
		stdout(p)
		os.Exit(1)
	}

	given := os.Args[1]

	if given == "" {
		p.Err = `Given empty string, please provide a valid statement`
	}

	stmt, err := sqlparser.Parse(given)

	if err != nil {
		p.Err = err.Error()
		stdout(p)
		os.Exit(1)
	}

	beep, err := json.Marshal(stmt)
	if err != nil {
		fmt.Println(err)
	}

	switch stmt.(type) {
	case *sqlparser.DDL:
		stdout(stmt)
	case *sqlparser.Select:
		p.Data = string(beep)
		stdout(p)
	case *sqlparser.Insert:
		p.Data = string(beep)
		stdout(p)
	case *sqlparser.Update:
		p.Data = string(beep)
		stdout(p)
	case *sqlparser.Delete:
		p.Data = string(beep)
		stdout(p)
	case *sqlparser.CreateTable:
		p.Data = string(beep)
		stdout(p)
	}
}

func stdout(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	os.Exit(0)
}
