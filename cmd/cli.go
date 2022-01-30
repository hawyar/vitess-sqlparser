package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"os"
)

type Payload struct {
	Type string `json:"type"` // type of stmt given
	Data string `json:"data"`
	Err  string `json:"err"`
}

func main() {

	p := Payload{}

	flag.Parse()

	if len(os.Args) < 2 {
		p.Err = "empty-query"
		stdout(p)
		os.Exit(1)
	}

	given := os.Args[1]

	if given == "" {
		p.Err = `Given empty string, please provide a valid statement`
	}

	tree, err := sqlparser.Parse(given)

	if err != nil {
		p.Err = err.Error()
		stdout(p)
		os.Exit(1)
	}
	stdout(tree)
}

func stdout(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	os.Exit(0)
}
