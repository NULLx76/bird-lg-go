package proxy

import (
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"strings"
)

type SummaryTable []PeerRow

type PeerRow struct {
	Name  string `json:"name"`
	Proto string `json:"proto"`
	Table string `json:"table"`
	State string `json:"state"`
	Since string `json:"since"`
	Info  string `json:"info"`
}

func (SummaryTable) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

type PeerDetails struct {
	Info    PeerRow `json:"info"`
	Details string  `json:"details"`
}

func (PeerDetails) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

type RouteDetails struct {
	Address string `json:"address"`
	Details string `json:"details"`
}

func (RouteDetails) Render(http.ResponseWriter, *http.Request) error {
	return nil
}

var headerRegex = regexp.MustCompile(`Name\s+Proto\s+Table\s+State\s+Since\s+Info`)
var columnRegex = regexp.MustCompile(`(\w+)\s+(\w+)\s+([\w-]+)\s+(\w+)\s+([0-9\-.:]+)\s?(.*)`)

func parsePeerRow(line string) PeerRow {
	split := columnRegex.FindStringSubmatch(line)

	// split[0] == whole string
	return PeerRow{
		Name:  strings.TrimSpace(split[1]),
		Proto: strings.TrimSpace(split[2]),
		Table: strings.TrimSpace(split[3]),
		State: strings.TrimSpace(split[4]),
		Since: strings.TrimSpace(split[5]),
		Info:  strings.TrimSpace(split[6]),
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
		Info:    header,
		Details: details[2],
	}, nil
}
