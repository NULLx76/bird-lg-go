package main

import (
	"errors"
	"regexp"
	"strings"
)

type SummaryTable []PeerRow

type PeerRow struct {
	name  string
	proto string
	table string
	state string
	since string
	info  string
}

type PeerDetails struct {
	header  PeerRow
	details string
}

var headerRegex = regexp.MustCompile(`Name\s+Proto\s+Table\s+State\s+Since\s+Info`)
var columnRegex = regexp.MustCompile(`(\w+)\s+(\w+)\s+([\w-]+)\s+(\w+)\s+([0-9\-]+)\s?(.*)`)

func parsePeerRow(line string) PeerRow {
	split := columnRegex.FindStringSubmatch(line)

	// split[0] == whole string
	return PeerRow{
		name:  strings.TrimSpace(split[1]),
		proto: strings.TrimSpace(split[2]),
		table: strings.TrimSpace(split[3]),
		state: strings.TrimSpace(split[4]),
		since: strings.TrimSpace(split[5]),
		info:  strings.TrimSpace(split[6]),
	}
}

func parseSummaryTable(str string) (SummaryTable, error) {
	rows := strings.Split(str, "\n")
	if !headerRegex.MatchString(rows[0]) {
		return nil, errors.New("invalid protocol table")
	}

	var table SummaryTable
	for i := 0; i < len(rows)-1; i++ {
		row := strings.TrimSpace(rows[i+1])
		if row == "" {
			continue
		}
		table = append(table, parsePeerRow(row))
	}

	return table, nil
}

func parsePeerDetails(str string) (*PeerDetails, error) {
	details := strings.SplitN(str, "\n", 3)
	if len(details) != 3 || !headerRegex.MatchString(details[0]) {
		return nil, errors.New("invalid protocol table")
	}

	details[1] = strings.TrimSpace(details[1])
	header := parsePeerRow(details[1])

	return &PeerDetails{
		header:  header,
		details: details[2],
	}, nil
}
