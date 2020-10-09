package main

import (
	"errors"
	"regexp"
	"strings"
)

type ProtocolTable []ProtocolRow

type ProtocolRow struct {
	name  string
	proto string
	table string
	state string
	since string
	info  string
}

var headerRegex = regexp.MustCompile(`Name\s+Proto\s+Table\s+State\s+Since\s+Info`)
var columnRegex = regexp.MustCompile(`(\w+)\s+(\w+)\s+([\w-]+)\s+(\w+)\s+([0-9\-]+)\s?(.*)`)

func parseProtocolRow(line string) ProtocolRow {
	split := columnRegex.FindStringSubmatch(line)

	// split[0] == whole string
	return ProtocolRow{
		name:  strings.TrimSpace(split[1]),
		proto: strings.TrimSpace(split[2]),
		table: strings.TrimSpace(split[3]),
		state: strings.TrimSpace(split[4]),
		since: strings.TrimSpace(split[5]),
		info:  strings.TrimSpace(split[6]),
	}
}

func parseProtocolTable(str string) (ProtocolTable, error) {
	rows := strings.Split(str, "\n")
	if !headerRegex.MatchString(rows[0]) {
		return nil, errors.New("invalid protocol table")
	}

	var table ProtocolTable
	for i := 0; i < len(rows)-1; i++ {
		row := strings.TrimSpace(rows[i+1])
		if row == "" {
			continue
		}
		table = append(table, parseProtocolRow(row))
	}

	return table, nil
}
